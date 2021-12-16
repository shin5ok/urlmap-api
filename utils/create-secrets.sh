#!/bin/bash
echo -n $DBUSER | gcloud secrets create DBUSER --data-file=-
echo -n $DBPASS | gcloud secrets create DBPASS --data-file=-
echo -n $DBNAME | gcloud secrets create DBNAME --data-file=-
echo -n $DBHOST | gcloud secrets create DBHOST --data-file=-

