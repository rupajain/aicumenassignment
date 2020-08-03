Functionality:

/add: Adds an employee specifying the name(compulsory), department(optional), address and skills (passed as request body parameters). An error response should be shown if the name is not provided. A unique ID should be auto-generated and assigned to each employee. 
 

/update: Updates the department, address or skills of the employee based on the request body parameters passed. 
 

/search: Lists details (ID, department, address and skills) of all employees whose name, department, address or skills matches the search term parameter specified in the request body. 
 

/list: Lists all employees if no parameters are passed. Employees can also be listed by ID, name or department if these specific parameters are passed. 
 

/delete: Deletes an employee. A deleted employee should only be deactivated unless the ‘permanentlyDelete’ parameter is passed. A deactivated employee should not be shown in /search, /list, . However, he/she can be activated again using /restore. An employee should be permanently deleted by calling /delete with the 'permanentlyDelete' parameter set to true. 
 

/restore: Restores a deactivated employee. 


Steps to run:
 1. download the source code .add to workspace (src pkg bin folders)copy the source code to src folder
 2. go to src folder and project folder then run  go install
 3. go to bin folder and run ./projectname it runs the project
 
