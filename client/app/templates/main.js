export const navbar = `
<nav class="navbar">
    <div class="navbar-left">
        <span class="navbar-title">Zawya</span>
    </div>
    <div class="navbar-right">
        <button id="logout-button" class="logout-button">Logout</button>
    </div>
</nav>
`;

export const postInputContainer = `
<div class="post-input-container">
    <input
        type="text"
        id="open-add-post"
        class="post-input"
        placeholder="What's on your mind?"
        readonly
    />
    <div class="post-container"></div>
</div>
`;

export const sidebar = `
<div class="sidebar">
    <div class="sidebar-header">
        <h2>Users</h2>
    </div>
    <div class="users-list">
    </div>
</div>
`;

export const commentsHeader = `
<div class="comments-header">
    <h3>Comments</h3>
    <button id="close-comments" class="close-button">×</button>
</div>
`;

export const commentsContainer = `
<div class="comments-container">
</div>
`;

export const commentInputContainer = `
<div class="comment-input-container">
    <input id="comment-input" type="text" placeholder="Write a comment..." />
    <button id="submit-comment">Send</button>
</div>
`;

export const categoryOptions = `
<div id="category" class="category-options">
    <div class="checkbox">tech</div>
    <div class="checkbox">sports</div>
    <div class="checkbox">games</div>
</div>
`;

export const messageHeader = `
<div class="message-header">
    <h2 id="user-nickname"></h2>
</div>
`;

export const messagesContainer = `
<div class="messages-container">
</div>
`;

export const messageInput = `
<div class="message-input">
    <input type="text" placeholder="Type your message here..." id="message"/>
    <button id="btn-message">Send</button>
    <div class="typing-indicator hidden typing" id="typing-indicator">
    <span id="typing-user"></span> 
    <div class="dots">
        <span></span>
        <span></span>
        <span></span>
    </div>
    </div>
</div>
`;

export const mainContent = `
<div class="main-content">
    ${navbar}
    <div id="messages-display" class="hidden">
        ${messageHeader}
        ${messagesContainer}
        ${messageInput}
    </div>
</div>
`;

export const addPostSection = `
<div id="add-post" class="add-post-section">
    ${categoryOptions}
    <textarea id="postContent" class="post-content-input" placeholder="Write your post here..." maxlength="500"></textarea>
    <div class="character-counter">0/500</div>
    <button class="button post-button" id="Posting">Post</button>
    <button id="close-add-post" class="close-button">Close</button>
</div>
`;

export const popupOverlay = `
<div id="popup-overlay" class="popup-overlay hidden">
    ${addPostSection}
</div>
`;

export const commentsPopup = `
<div class="comments-popup">
    ${commentsHeader}
    ${commentsContainer}
    ${commentInputContainer}
</div>
`;

export const commentsOverlay = `
<div id="comments-overlay" class="comments-overlay hidden">
    ${commentsPopup}
</div>
`;

export const appContainer = `
<div class="app-container">
    ${sidebar}
    ${mainContent}
    ${popupOverlay}
    ${commentsOverlay}
</div>
`;

export const main = `
${postInputContainer}
${appContainer}
`;
