Overview
- Basic of computer
- How programming work with computer
- Learn Golang, Nodejs
- Workshop create OWN elk stack
- Build web application (database design)

Today 8/9/2019
- Overview basic of computer & programming
- Overview Workshop project

Overview of web application

1. key Url (input) === request to server ===> 
2. receieve request and processing (process)  (read data from database)
3. send response (output) eg. <html>

School web application
1. student data
2. class data
3. student register to class
4. student schedule class
5. grade of student

log of school application
1. response status

database design

## Students
- id (1,2,3,4)
- name
- student_id
- image
- date_of_birth
- username
- password
- created_date
- updated_date

## Classes
- id
- subject_id
- name
- created_date
- updated_date

## Class Course
- id
- class_id
- start_date
- end_date
- created_date
- updated_date

class => class_schedule  (1 => many)
## Class Schedule
- id
- class_id 
- day
- time_start
- time_end
- created_date
- updated_date

## Registration
- id
- student_id
- course_id
- grade
- point
- created_date
- updated_date
student => register <= course
1 => M <= 1 (many to many)

Data example
### Classes
id  name
1   thai language
### Class course
id  class_id  start_date  end_date
1   1         1/5/2019    1/11/2019
### Class schedule
id  course_id day     time_start  time_end  
1   1         Monday  10:00       12:00     
2   1         Tuesday 9:00        11:00


# Workshop Day2
- design database 

- collect employee detail
- separate by employee section
- collect Leave information of employee
