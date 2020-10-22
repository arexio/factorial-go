package main

var indexTemplate = `<p>
<h1>Sample Server for the Factorial API</h1>
<a href="/auth/factorial">
	<img src="/static/factorial_logo.png" alt="ConnectToFactorial">
	<p>Connect To Factorial</p>
</a>
</p>`

var connectedTemplate = `
<h1>Sample Server for the Factorial API</h1>
<p>AccessToken: {{.AccessToken}}</p>
<p>TokenType: {{.TokenType}}</p>
<p>RefreshToken: {{.RefreshToken}}</p>
<p>Expiry: {{.Expiry}}</p>

<p><a href="/employees"/>Employees</p>
<p><a href="/folders"/>Folders</p>
<p><a href="/leave_types"/>Leave Types</p>
<p><a href="/leaves"/>Leaves</p>`

var employeesTemplate = `
<h1>Employees</h1>
{{range .Employees}}
<p>--  <b>ID:</b>{{.ID}}  |  <b>BirthdayOn:</b>{{.BirthdayOn}}  |  <b>StartDate:</b>{{.StartDate}}</p>
<p>--  <b>Email:</b>{{.Email}}  |  <b>FullName:</b>{{.FullName}}  |  <b>FirstName:</b>{{.FirstName}}</p>
<p>--  <b>LastName:</b>{{.LastName}}  |  <b>ManagerID:</b>{{.ManagerID}}  |  <b>Role:</b>{{.Role}}</p>
<p>--  <b>TimeoffManagerID:</b>{{.TimeoffManagerID}}  |  <b>TerminatedOn:</b>{{.TerminatedOn}}  |  <b>PhoneNumber:</b>{{.PhoneNumber}}</p>
<p>--  <b>Gender:</b>{{.Gender}}  |  <b>Nationality:</b>{{.Nationality}}  |  <b>BankNumber:</b>{{.BankNumber}}</p>
<p>--  <b>Country:</b>{{.Country}}  |  <b>City:</b>{{.City}}  |  <b>State:</b>{{.State}}</p>
<p>--  <b>PostalCode:</b>{{.PostalCode}}  |  <b>AddresLine1:</b>{{.AddresLine1}}  |  <b>AddressLine2:</b>{{.AddressLine2}}</p>
<p>--  <b>SocialSecurityNumber:</b>{{.SocialSecurityNumber}}  |  <b>CompanyHolidayIDs:</b>{{.CompanyHolidayIDs}}  |  <b>Identifier:</b>{{.Identifier}}</p>
<p>--  <b>IdentifierType:</b>{{.IdentifierType}}  |  <b>Hiring:</b>{{.Hiring}}  |  <b>LocationID:</b>{{.LocationID}}</p>
<p>--  <b>TeamIDs:</b>{{.TeamIDs}}</p>
{{end}}
`

var foldersTemplate = `
<h1>Folders</h1>
{{range .Folders}}
<p>--  <b>ID:</b>{{.ID}}  |  <b>CompanyID:</b>{{.CompanyID}}  |  <b>Name:</b>{{.Name}}</p>
<p>--  <b>Type:</b>{{.Type}}  |  <b>Active:</b>{{.Active}}  |  <b>CreatedAt:</b>{{.CreatedAt}}</p>
<p>--  <b>UpdatedAt:</b>{{.UpdatedAt}}
{{end}}
`
var leaveTypesTemplate = `
<h1>Leave Types</h1>
{{range .LeaveTypes}}
<p>--  <b>ID:</b>{{.ID}}  |  <b>Accrues:</b>{{.Accrues}}  |  <b>Active:</b>{{.Active}}</p>
<p>--  <b>ApprovalRquired:</b>{{.ApprovalRquired}}  |  <b>Attachment:</b>{{.Attachment}}  |  <b>Color:</b>{{.Color}}</p>
<p>--  <b>Identifier:</b>{{.Identifier}} |  <b>Name:</b>{{.Name}}  |  <b>Visibility:</b>{{.Visibility}}</p>
<p>--  <b>Workable:</b>{{.Workable}}
{{end}}
`

var leavesTemplate = `
<h1>Leaves</h1>
{{range .Leaves}}
<p>--  <b>ID:</b>{{.ID}}  |  <b>Description:</b>{{.Description}}  |  <b>EmployeeID:</b>{{.EmployeeID}}</p>
<p>--  <b>FinishOn:</b>{{.FinishOn}}  |  <b>HalfDay:</b>{{.HalfDay}}  |  <b>LeaveTypeID:</b>{{.LeaveTypeID}}</p>
<p>--  <b>StartOn:</b>{{.StartOn}}
{{end}}
`
