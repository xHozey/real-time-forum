export const main = `

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

<div class="app-container">
    <div class="sidebar">
        <div class="sidebar-header">
            <h2>Users</h2>
        </div>
        <div class="users-list">
        </div>
    </div>

    <div class="main-content">
      <nav class="navbar">
    <div class="navbar-left">
      <span class="navbar-title">Zawya</span>
    </div>
    <div class="navbar-right">
      <button id="logout-button" class="logout-button">Logout</button>
    </div>
  </nav>

        <div id="messages-display" class="hidden">
            <div class="message-header">
                <h2 id="user-nickname"></h2>
            </div>
            <div class="messages-container">
            </div>
            <div class="message-input">
                <input type="text" placeholder="Type your message here..." id="message"/>
                <button id="btn-message">Send</button>
            </div>
        </div>
    </div>

    <div id="popup-overlay" class="popup-overlay hidden">
        <div id="add-post" class="add-post-section">
            <div id="category" class="category-options">
                <div class="checkbox">tech</div>
                <div class="checkbox">sports</div>
                <div class="checkbox">games</div>
            </div>
            <textarea id="postContent" class="post-content-input" placeholder="Write your post here..." maxlength="500"></textarea>
            <div class="character-counter">0/500</div>
            <button class="button post-button" id="Posting">Post</button>
            <button id="close-add-post" class="close-button">Close</button>
        </div>
    </div>
    
<div id="comments-overlay" class="comments-overlay hidden">
    <div class="comments-popup">
        <div class="comments-header">
            <h3>Comments</h3>
            <button id="close-comments" class="close-button">×</button>
        </div>
        <div class="comments-container">
        </div>
        <div class="comment-input-container">
        <input id="comment-input" type="text" placeholder="Write a comment..." />
        <button id="submit-comment">Send</button>
        </div>
</div>
    
</div>
`;
