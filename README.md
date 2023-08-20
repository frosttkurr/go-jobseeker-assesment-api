Documentation
API CANDIDATES - Jobseeker Assessment Test
This program is used to create API for candidates data using Go with Gin framework.

Before running the project, you should install needed libraries following command:

go get
Then setup port and database configuration on .env (copy from .env.example). Example:

PORT=3004
DB_URL="root:@tcp(localhost:3310)/db_candidate"
Finally, you can try to run to start API.

go run main.go
﻿

This is the example of successfuly working API:


﻿
Note:

Database for the program is inside of database folder.

GET
Get All Candidates
http://localhost:3004/api/candidates
You can get a list of candidates data alongside pagination and filtering for server-side datatables.

For example:

{{baseUrl}}/api/candidates?page=1&pageSize=5&gender=M 
Description:
page used to specify page number for pagination

pageSize used to sets the number of results to be displayed per page

gender used to filter data based on candidates gender

List of available filter parameters:
email : filter based on email

phoneNumber : filter based on phone number

dob : filter based on date of birth

pob : filter based on place of birth

gender : filter based on gender

yearExp : filter based on years experience

lastSalary : filter based on last salary

﻿

GET
Show Candidate Based on ID
http://localhost:3004/api/candidates/2
You can show candidate data based on its ID

For example:

{{baseUrl}}/api/candidates/16
﻿

POST
Insert New Candidate
http://localhost:3004/api/candidates
You can create new candidate data using this endpoint

﻿

Body
raw (json)
json
{
    "full_name": "John Cena",
    "email": "email99@example.com",
    "phone_number": "33333333333",
    "dob": "2000-03-12",
    "pob": "City7",
    "gender": "F",
    "year_exp": "2 - 3 years",
    "last_salary": "Rp4.000.000"
}
PUT
Update Candidate
http://localhost:3004/api/candidates/1
You can update existing candidate data using based on its ID

For example:

{{baseUrl}}/api/candidates/15
﻿

Body
raw (json)
json
{
    "email": "email222@example.com",
    "phone_number": "987654321340",
    "full_name": "Jane Smith",
    "dob": "1995-02-15",
    "pob": "City2",
    "gender": "F",
    "year_exp": "2 - 3 years",
    "last_salary": "Rp4.000.000"
}
DELETE
Delete Candidate
http://localhost:3004/api/candidates/1
You can delete existing candidate data using based on its ID

For example:

{{baseUrl}}/api/candidates/15
﻿
