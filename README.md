# MailMind API

**MailMind Backend** is the server-side application that powers the MailMind Chrome extension. Built with Golang and the Gin framework, it provides authentication, email processing, and AI-driven responses for seamless email management.

**Note:** This backend is still under development. Stay tuned for updates!

## Features  

* **OAuth2 Authentication** – Secure Gmail login and token handling.  
* **AI-Powered Email Replies** – Generates smart email responses using Gemini AI.  
* **Draft Management** – Save and retrieve email drafts efficiently.  
* **Email Sending** – Send emails via Gmail’s API.  
* ⚙**Customizable AI Settings** – Users can configure AI-generated reply preferences.  
* **Logging & Error Handling** – Structured logging for better debugging and monitoring.  

## Technologies  

* **Golang** – High-performance backend language.  
* **Gin** – Lightweight and fast web framework for handling API requests.  
* **MongoDB** – NoSQL database for storing user preferences, drafts, and AI-generated responses.  
* **OAuth2** – Secure authentication via Gmail.  
* **Gemini AI** – Google’s AI model for generating email replies.  

## API Endpoints  

### Authentication  
- `POST /auth/connect` – Redirects users to Gmail OAuth login.  
- `GET /auth/callback` – Handles OAuth callback and token exchange.  
- `GET /auth/logout` – Revokes access and logs the user out.  

### AI-Generated Replies  
- `POST /ai/generate-reply` – Generates an AI-powered response based on email content.  

### Email Management  
- `POST /emails/send` – Sends an email.  
- `POST /emails/draft` – Saves a draft email.

  and more will be added soon ...

## Getting Started  

### Prerequisites  

* Go 1.18+ installed  
* MongoDB instance running  
* Google Cloud API credentials (for OAuth2 & Gemini AI)  
* Gmail API enabled in Google Cloud Console  

### Installation  

1. **Clone the repository:**  
   ```bash  
   git clone [your-repo-url]  
   cd mailmind-backend  
   ```  

2. **Install dependencies:**  
   ```bash  
   go mod tidy  
   ```  

3. **Set up environment variables:**  
   Create a `.env` file and configure:  
   ```env  
   ACCESS_TOKEN_SECRET=your-secret-key  
   MONGODB_URI=mongodb://localhost:27017/mailmind  
   GMAIL_CLIENT_ID=your-client-id  
   GMAIL_CLIENT_SECRET=your-client-secret  
   GMAIL_REDIRECT_URI=your-redirect-uri  
   ```  

4. **Run the server:**  
   ```bash  
   go run main.go  
   ```  

5. **Test the API with Postman or Curl:**  
   ```bash  
   curl -X POST http://localhost:8080/v1/ai/generate-reply -d '{"email_id": "some_id", "content": "Hello, how are you?"}'  
   ```  

## Contributing  

We welcome contributions to improve **MailMind Backend**! Feel free to submit a pull request or report issues.  

## License  

This project is licensed under the **raufzer**.  

