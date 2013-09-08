uploadthis
==========

A flexible, performant command line S3 uploader by your friends at ShareThis


Goal
====
To provide a process that can upload log data to s3 that is easily configurable through a YAML file.  The process should be flexible enough to handle many different use-cases such as time-based or size-based uploads while still being performant.

Problem Description
===================
Log analysis is one of the most effective way to do analysis at a large scale.  As a first step in most analytic pipelines it is required to store and backup files in the cloud.  This however brings many complexities like time-based or size-based rollovers, bandwidth constraints, file-sizes, atomic operations and most importantly performance.  The burden is placed on the developer to handle all these operational complexities.  There should be a better way.

Example Usage
==============
uploadThis -c /etc/uploadconf.yaml

Building
=========

How to build the code:

  go get github.com/imosquera/uploadthis
  cd $GOPATH/src/github.com/imosquera/uploadthis
  go build cmd/uploadthis.go


 In short, the quickest way to get going with annotator is to include the following in the `<head>` of your document (paths relative to the repository root):

    <script src='lib/vendor/jquery.js'></script>
    <script src='lib/vendor/json2.js'></script>
    go get github.com/imosquera/uploadthis
    cd $GOPATH/src/github.com/imosquera/uploadthis
    go build cmd/uploadthis.go

    <script src='pkg/annotator.min.js'></script>
    <link rel='stylesheet' href='pkg/annotator.min.css'>
