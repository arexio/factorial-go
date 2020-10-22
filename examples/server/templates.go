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
<p><a href="/leaves"/>Leaves</p>
<p><a href="/webhooks"/>Webhooks</p>
<p><a href="/documents"/>Documents</p>
<p><a href="/hiring_versions"/>Hiring Versions</p>
<p><a href="/company_holidays"/>Company Holidays</p>`

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
<p>--  <b>UpdatedAt:</b>{{.UpdatedAt}}</p>
{{end}}
`
var leaveTypesTemplate = `
<h1>Leave Types</h1>
{{range .LeaveTypes}}
<p>--  <b>ID:</b>{{.ID}}  |  <b>Accrues:</b>{{.Accrues}}  |  <b>Active:</b>{{.Active}}</p>
<p>--  <b>ApprovalRquired:</b>{{.ApprovalRquired}}  |  <b>Attachment:</b>{{.Attachment}}  |  <b>Color:</b>{{.Color}}</p>
<p>--  <b>Identifier:</b>{{.Identifier}} |  <b>Name:</b>{{.Name}}  |  <b>Visibility:</b>{{.Visibility}}</p>
<p>--  <b>Workable:</b>{{.Workable}}</p>
{{end}}
`

var leavesTemplate = `
<h1>Leaves</h1>
{{range .Leaves}}
<p>--  <b>ID:</b>{{.ID}}  |  <b>Description:</b>{{.Description}}  |  <b>EmployeeID:</b>{{.EmployeeID}}</p>
<p>--  <b>FinishOn:</b>{{.FinishOn}}  |  <b>HalfDay:</b>{{.HalfDay}}  |  <b>LeaveTypeID:</b>{{.LeaveTypeID}}</p>
<p>--  <b>StartOn:</b>{{.StartOn}}</p>
{{end}}
`
var webhooksTemplate = `
<h1>Webhooks</h1>
{{range .Webhooks}}
<p>--  <b>SubscriptionType:</b>{{.SubscriptionType}}</p>
{{end}}
`

var documentsTemplate = `
<h1>Documents</h1>
{{range .Documents}}
<p>--  <b>ID:</b>{{.ID}}  |  <b>EmployeeID:</b>{{.EmployeeID}}  |  <b>CompanyID:</b>{{.CompanyID}}</p>
<p>--  <b>FolderID:</b>{{.FolderID}}  |  <b>File:</b>{{.File}}  |  <b>FileName:</b>{{.FileName}}</p>
<p>--  <b>Public:</b>{{.Public}} | <b>CreatedAt:</b>{{.CreatedAt}}  |  <b>UpdatedAt:</b>{{.UpdatedAt}}</p>
{{end}}
`

var hiringVersionsTemplate = `
<h1>Hiring Versions</h1>
{{range .HiringVersions}}
<p>--  <b>ID:</b>{{.ID}}  |  <b>EffectiveOn:</b>{{.EffectiveOn}}  |  <b>EmployeeID:</b>{{.EmployeeID}}</p>
<p>--  <b>BaseCompensationAmountInCents:</b>{{.BaseCompensationAmountInCents}}  |  <b>BaseCompensationType:</b>{{.BaseCompensationType}}  |  <b>StartDate:</b>{{.StartDate}}</p>
<p>--  <b>EndDate:</b>{{.EndDate}} | <b>JobTitle:</b>{{.JobTitle}}  |  <b>WorkingHoursInCents:</b>{{.WorkingHoursInCents}}</p>
<p>--  <b>WorkingPeriodUnit:</b>{{.WorkingPeriodUnit}}</p>`

var companyHolidaysTemplate = `
<h1>Company Holidays</h1>
{{range .CompanyHolidays}}
<p>--  <b>ID:</b>{{.ID}}  |  <b>Summary:</b>{{.Summary}}  |  <b>Description:</b>{{.Description}}</p>
<p>--  <b>Date:</b>{{.Date}}  |  <b>HalfDay:</b>{{.HalfDay}}  |  <b>LocationID:</b>{{.LocationID}}</p>
{{end}}
`
