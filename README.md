# Advanced Forum with Real-Time Private Messaging

## Project Overview

This project is an upgraded forum application featuring:

- User registration and login  
- Creation of posts with categories  
- Commenting on posts  
- Real-time private messaging between users  

The forum is designed as a **Single Page Application (SPA)** with one HTML file, where all navigation and page updates are handled via JavaScript.

---

## Features

### Registration and Login

- Users register with the following information:
  - Nickname  
  - Age  
  - Gender  
  - First Name  
  - Last Name  
  - E-mail  
  - Password  
- Users can log in using either their **nickname or e-mail** combined with their password.  
- Users can log out from any page in the forum.  

### Posts and Comments

- Users can create posts assigned to specific categories.  
- Users can comment on posts.  
- Posts appear in a feed display.  
- Comments are visible only after clicking on a post.  

### Private Messaging (Real-Time Chat)

- Displays a list of users who are online/offline and available to chat.  
- The user list is sorted by the last message sent (like Discord); new users without messages are sorted alphabetically.  
- Users can send private messages in real-time using WebSockets.  
- The chat window shows past messages:  
  - Loads the last 10 messages initially.  
  - Loads 10 more messages on scroll-up with throttling/debouncing to avoid spamming.  
- Messages show:  
  - Timestamp of when the message was sent  
  - Username of the sender  
- New messages are received instantly without refreshing the page (via WebSockets).  

---

## Technologies Used

- **Backend:** Go (Golang)  
- **Database:** SQLite  
- **Authentication:** bcrypt for password hashing  
- **Real-Time:** Gorilla WebSocket library for WebSocket connections  
- **UUID:** gofrs/uuid or google/uuid for unique identifiers  
- **Frontend:** Vanilla JavaScript, HTML, and CSS (no frontend frameworks or libraries allowed)  

---

## Allowed Packages

- Standard Go packages  
- Gorilla WebSocket  
- SQLite3 driver  
- bcrypt  
- UUID libraries (`gofrs/uuid` or `google/uuid`)  

---

## What You'll Learn

- Core web fundamentals: HTML, CSS, HTTP, Sessions, and Cookies  
- Frontend and backend communication via WebSockets  
- Handling databases and SQL with SQLite  
- Managing concurrency with Go routines and channels  
- DOM manipulation with vanilla JavaScript  
- Implementing real-time features in web applications  

---

## Usage

1. Run the Go backend server (which handles REST API and WebSocket connections).  
2. Open the forum in a web browser (served by the Go server).  
3. Register or log in to access forum features.  
4. Create posts, comment, and chat privately in real-time with other users.  

---

## Notes

- The application must be robust with proper error handling.  
- All navigation and dynamic content updates should be handled client-side with JavaScript.  
- The project prohibits using frontend frameworks (React, Angular, Vue, etc.).  

---

## License

This project is for educational and learning purposes.
