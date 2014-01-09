from optparse import OptionParser
from boto.s3.connection import S3Connection
from boto.exception import S3ResponseError
from subprocess import call
import filecmp
import time
import os


def get_file(bucket_name, document, output, aws_id=None, aws_key=None):
    #http://ceph.com/docs/next/radosgw/s3/python/
    print "Connect to S3"
    #AWS credentials must be set as environment variables (AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY). Otherwise use S3Connection with parameters
    if aws_id and aws_key:
        conn = S3Connection(aws_id, aws_key)
    else:
        conn = S3Connection()

    try:
        print "Get bucket: {0}".format(bucket_name)
        bkt = conn.get_bucket(bucket_name)
        document = "{0}/{1}".format(time.strftime("%Y/%m/%d/%H", time.gmtime()), document)
        print "Get file key: {0}".format(document)
        key = bkt.get_key(document)
        if key is None:
            print '\033[31m S3 ERROR: Cannot get file key from S3 for path: {0} \033[0m'.format(document)
            exit(1)
        print "Download to: {0}".format(output)
        key.get_contents_to_filename(output)
        print "Delete document"
        bkt.delete_key(document)
    except S3ResponseError as e:
        print '\033[31m S3 ERROR: {0} \033[0m'.format(e)
        exit(1)


def check_file(original, downloaded):
    print "Unzip downloaded archive: {0}".format(downloaded)
    try:
        call(['gzip', '-d', downloaded])
    except OSError as e:
        return e
    print "Compare files: {0} == {1}".format(original, downloaded[:-3])
    if filecmp.cmp(original, downloaded[:-3]):
        return 0
    else:
        return "Files do not match"


if __name__ == '__main__':
    parser = OptionParser()
    parser.add_option("-b", "--bucket",   dest="bucket",    default="loopy-analytics", metavar="test",      help="S3 bucket name")
    parser.add_option("-d", "--document", dest="document",  default="test.json.gz",    metavar="doc.txt",   help="S3 document name")
    parser.add_option("-m", "--monDir",   dest="monitordir",default="/tmp/fixtures/",  metavar="/tmp/log/", help="Monitor directory")
    parser.add_option("-e", "--emptyDoc", dest="empty_doc", default="test.empty.json", metavar="empty.txt", help="Filename of local empty log that should not be processed")
    parser.add_option("-o", "--output",   dest="output",    default="/tmp/fixtures/download/test.json.gz", metavar="/tmp/downloaded.txt", help="output file location to save to")
    parser.add_option("-c", "--compare",  dest="compare",   default="/tmp/fixtures/test.json",         metavar="/tmp/original.txt",   help="original file to compare with")
    parser.add_option("-i", "--aws_id",   dest="aws_id",    default="",                metavar="MY KEY ID", help="AWS Key ID")
    parser.add_option("-a", "--aws_key",  dest="aws_key",   default="",                metavar="MY SECRET", help="AWS Secret Key")

    options, args = parser.parse_args()

    opts = {}
    for opt, value in options.__dict__.items():
        if value:
            opts[opt] = value

    if opts['empty_doc'] and not os.path.exists(opts['empty_doc']):
        print "\033[31m ERROR. Empty log file missed but had to be ignored\033[0m"
        exit(1)

    #Check uploader temp directories to be empty after files processing
    monitordir = opts['monitordir']
    if not monitordir.endswith(os.path.sep):
        monitordir += os.path.sep
    tmp_dirs = ["compress", "upload"]
    for sub_dir in tmp_dirs:
        tmp_dir = "{0}.uploadthis/{1}".format(monitordir, sub_dir)
        if os.listdir(tmp_dir) != []:
            print "\033[31m ERROR. {0} folder is not empty after test\033[0m".format(tmp_dir)
            exit(1)

    if not opts['aws_id'] or not opts['aws_key']:
        print "AWS credentials not provided. Assume AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables are set"

    get_file(opts['bucket'], opts['document'], opts['output'], aws_id=opts['aws_id'], aws_key=opts['aws_key'])

    err = check_file(opts['compare'], opts['output'])
    if not err:
        print "\033[32m TEST PASS\033[0m"
    else:
        print "\033[31m ERROR. {0}\033[0m".format(err)
        exit(1)
