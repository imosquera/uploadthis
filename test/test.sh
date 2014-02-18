#!/bin/bash

usage() {
    cat << EOF
    Integration test for S3 uploader
    usage: $0 path/to/python_test.py <AWS_KEY_ID> <AWS_ACCESS_SECRET>
EOF
}

if [ "$#" -ne 3 ]; then
  usage
  exit 1
fi

FIXTURES=${GOPATH}/src/github.com/imosquera/uploadthis/fixtures
TEST_FILE=test.json
LOGS=${FIXTURES}/monitordir/${TEST_FILE}
CONFIG_SAMPLE=${FIXTURES}/sample-config.yaml

TEMP_PATH=/tmp/fixtures
LOG_PATH=${TEMP_PATH}/monitordir
CONFIG=${TEMP_PATH}/config.yaml
LOGGING=/tmp/
TESTPATH=${TEMP_PATH}/download
OLD_ARCHIVED_LOG=${LOG_PATH}/.uploadthis/archive/event.old.log
S3_BUCKET=loopy-analytics

echo "Create temp directory"
mkdir -p ${LOG_PATH}
mkdir -p ${LOGGING}
mkdir -p ${TESTPATH}
mkdir -p ${LOG_PATH}/.uploadthis/archive

echo "Copy files"
cp ${LOGS} ${LOG_PATH}
cp ${LOGS} ${TEMP_PATH}
touch ${LOG_PATH}/test.empty.json
echo "Very old log file to be deleted from archive" > ${OLD_ARCHIVED_LOG}
touch -amt 201310010000 ${OLD_ARCHIVED_LOG}

echo "Setting up config parameters"
sed -e "s/myaccesskey/$2/" -e "s/mysupersecretkey/$3/" -e "s/bucket: 34/bucket: ${S3_BUCKET}/" -e "s|/tmp/logs/loopy/event/|${LOG_PATH}|" -e "s|/var/log/|${LOGGING}|" < ${CONFIG_SAMPLE} > ${CONFIG}

echo "Run Upload tool"
cd ${GOPATH}
bin/uploadthis -c ${CONFIG}

echo "Check archive cleanup"
ARCHIVED_FILES_NUMBER=$(ls -l $(dirname ${OLD_ARCHIVED_LOG}) | wc -l)
if [ ${ARCHIVED_FILES_NUMBER} != 2 ]
then
    echo "Archive cleanup test failed"
    exit 1
fi

echo "Check S3 content"
echo python $1 -i $2 -a $3 -b ${S3_BUCKET} -d ${TEST_FILE}.gz -m ${LOG_PATH} -o ${TESTPATH}/${TEST_FILE}.gz -c ${TEMP_PATH}/${TEST_FILE} -e ${LOG_PATH}/test.empty.json
python $1 -i $2 -a $3 -b ${S3_BUCKET} -d ${TEST_FILE}.gz -m ${LOG_PATH} -o ${TESTPATH}/${TEST_FILE}.gz -c ${TEMP_PATH}/${TEST_FILE} -e ${LOG_PATH}/test.empty.json

echo "Cleanup"
rm -r ${TEMP_PATH}
echo "Done"
