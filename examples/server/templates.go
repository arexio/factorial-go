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
<p><a href="/folders"/>Folders</p>`

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
<p>--  <b>SocialSecurityNumber:</b>{{.SocialSecurityNumber}}  |  <b>CompanyHolidaysIDs:</b>{{.CompanyHolidaysIDs}}  |  <b>Identifier:</b>{{.Identifier}}</p>
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
