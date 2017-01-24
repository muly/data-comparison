# data-comparison
go app to compare data in one flat file to data in another flat file, join the records by using the specified key fields, and report the changed records and fields.

## challenges
The challenge is that there are different data sets that needs to be tested. dataset A to dataset A, and dataset B to dataset B, for example, customer dataset, invoice dataset, leads dataset. And there is no fixed variants of data sets. So application cannot define the structs at compile time.

Another challenge is that, there can be n number of dataset of type A, that needs to be compared to find the delta. For example, 10 flat files of customer dataset that are generated in last 10 days, one each day , and we need to find how the customer data is changing over time.
