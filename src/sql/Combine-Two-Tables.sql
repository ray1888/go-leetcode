select Person.FirstName, Person.LastName, Address.City, Address.State from Person
Left Join Address on Person.PersonID = Address.PersonID