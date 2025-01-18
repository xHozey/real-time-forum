export const main = `<div class="app-container">
    <div class="sidebar">
        <div class="sidebar-header">
            <h2>Users</h2>
        </div>
        <div class="users-list">
        </div>
    </div>

    <div class="main-content">
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
    <div id="posts"></div>
</div>
`;
