# event-manager-crud-api
A simple and lightweight crud api for creating and managing events implemented using golang and gin. SQLite is used as the DB and endpoints are protected using JWT token.

The various endpoints are : 
1. GET /getEvents : Gets a list of events present in the DB.
2. GET /getEvent/:id : Gets a particular event according to the event id passed in request parameter
3. POST /createEvent : Create and saves an event entry in DB, only logged in users can create event. Needs JWT token to be passed in authorization header.
4. PUT /event/:id :  Updates a particular event entry, only only logged in and creator of the event can update the event. Needs JWT token to be passed in authorization header.
5. DELETE /deleteEvent/:id : Deletes a particular event entry, only only logged in and creator of the event can delete the event. Needs JWT token to be passed in authorization header.
6. POST /createUser : Signs up user
7. POST /login : Logs the user in after successfully authentication and return the jwt token to passed for subsequent request.
