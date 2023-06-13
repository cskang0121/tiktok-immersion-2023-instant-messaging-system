-- View all messages
SELECT * FROM messages; 

-- View all messages based on 
-- 1. chat 
-- 2. cursor
-- 3. limit
-- 4. reverse 
-- provided in the pull request
SELECT * FROM messages WHERE chat = 'a1:a2' ORDER BY send_time ASC LIMIT 2, 2;

-- View the total number of messages of a specific chat
SELECT COUNT(id) FROM messages WHERE chat = 'a1:a2';

-- Define the table "messages"
CREATE TABLE messages (id INT PRIMARY KEY AUTO_INCREMENT, chat VARCHAR(255), sender VARCHAR(255), send_time INT, message TEXT);

-- Test inserting into the table
INSERT INTO messages (chat, sender, send_time, message) VALUES ('a1:a2', 'a1', unix_timestamp(now()), 'hello');
INSERT INTO messages (chat, sender, send_time, message) VALUES ('a1:a2', 'a2', unix_timestamp(now()), 'hi there');
INSERT INTO messages (chat, sender, send_time, message) VALUES ('a1:a2', 'a1', unix_timestamp(now()), 'how is life?');
INSERT INTO messages (chat, sender, send_time, message) VALUES ('a1:a2', 'a2', unix_timestamp(now()), 'everything is good. how about you?');
INSERT INTO messages (chat, sender, send_time, message) VALUES ('a1:a2', 'a2', unix_timestamp(now()), 'same... wanna hangout?');

-- Test dropping the table
DROP TABLE messages;

-- Test deleting the table 
DELETE FROM messages;