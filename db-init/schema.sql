CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	"name" varchar(50) NULL
);

CREATE TABLE IF NOT EXISTS rides (
    id SERIAL PRIMARY KEY,          
    source TEXT,   
    destination TEXT, 
    distance INT,          
    cost INT               
);

CREATE TABLE IF NOT EXISTS bookings (
    id SERIAL PRIMARY KEY,           
    user_id INT NOT NULL,            
    ride_id INT NOT NULL,            
    time TIMESTAMP NOT NULL,         
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_ride FOREIGN KEY (ride_id) REFERENCES rides(id)
);