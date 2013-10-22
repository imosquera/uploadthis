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

FIXTURES=$GOPATH/src/github.com/imosquera/uploadthis/fixtures
TEST_FILE=test.json
LOGS=$FIXTURES/monitordir/$TEST_FILE
CONFIG_SAMPLE=$FIXTURES/sample-config.yaml

TEMP_PATH=/tmp/fixtures
LOG_PATH=$TEMP_PATH/monitordir
CONFIG=$TEMP_PATH/config.yaml
LOGGING=$TEMP_PATH/log
TESTPATH=$TEMP_PATH/download

S3_BUCKET=loopy-analytics

echo "Create temp directory"
mkdir -p $LOG_PATH
mkdir -p $LOGGING
mkdir -p $TESTPATH

echo "Copy files"
cp $LOGS $LOG_PATH
cp $LOGS $TEMP_PATH
echo "Update config"
sed -e "s/myaccesskey/$2/" -e "s/mysupersecretkey/$3/" -e "s/bucket: 34/bucket: $S3_BUCKET/" -e "s|/var/log/|$LOGGING|" < $CONFIG_SAMPLE > $CONFIG

echo "Run Upload tool"
cd $GOPATH
bin/uploadthis -c $CONFIG

echo "Check upload"
echo python $1 -i $2 -a $3 -b $S3_BUCKET -d $TEST_FILE.gz -o $TESTPATH/$TEST_FILE.gz -c $TEMP_PATH/$TEST_FILE
python $1 -i $2 -a $3 -b $S3_BUCKET -d $TEST_FILE.gz -o $TESTPATH/$TEST_FILE.gz -c $TEMP_PATH/$TEST_FILE

echo "Cleanup"
rm -r $TEMP_PATH
echo "Done"

