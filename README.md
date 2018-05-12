# Sokoban server

Backend for kaj-sokoban - provides the REST API which the client uses for obtaining game maps, scoreboards and posting scores. 

## Setup
 1. Compile the go code
 2. Get a mysql database and create required tables with the provided script (```database.sql```)
 3. Setup the database connection in the ```settings.ini``` file
 4. Make sure the compiled binary and ```settings.ini``` are in the same folder
 5. Run the compiled binary
 6. Add some maps to the database (table ```map```)

## Related projects
 * **[Sokoban client](https://github.com/Silaedru/kaj-sokoban-client)**
 * **[Sokoban mapeditor](https://github.com/Silaedru/kaj-sokoban-mapeditor)**