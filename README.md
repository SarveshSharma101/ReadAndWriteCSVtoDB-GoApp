# ReadAndWriteCSVtoDB-GoApp

## This is a project which has 2 go applicaiton in it:
1. Read and write file app
This app picks up '.csv' files from a folder, read them one by one and for each line of data in the csv file the app make a post request to an api which will save this read data.

2. Go Server to save incoming files data to my sql DB
This app is basically a server that is listening to the incoming post request, after getting a post request the app save the data to the DB.
