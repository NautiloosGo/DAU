# DAU

Takes data from .csv
  example:
  [date link dau]

Export to json file //in progress

Make .xlsx file with charts //to do




proccesses step by step:
1. download old data from internal json files (create new if not found)
2. find files in directory "originaldata" with name *DAU*.csv and *Partners*.csv
3. for each file start goroutine.
  - read by rows
  - edit links, dates and dau as we want to
  - push each row in channel
  - exit by EOF
4. start goroutines to collect data from channels
  - skip duplicates
  - sum results for same date and partner with different links //only for partners
  - ? stop goroutines
5. upload data to json when all goroutines stoped //PROBLEM
6. create xlsx file by requirements //todo
