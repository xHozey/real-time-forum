export const main = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Discord-like UI</title>
    <link rel="stylesheet" href="styles.css">
</head>
<body>
    <div class="container">
        <!-- Left Sidebar -->
        <div class="sidebar">
            <div class="sidebar-header">
                <h2>Servers</h2>
            </div>
            <div class="server-list">
                <div class="server-item">Server 1</div>
                <div class="server-item">Server 2</div>
                <div class="server-item">Server 3</div>
            </div>
            <div class="sidebar-footer">
                <button class="add-server-btn">+</button>
            </div>
        </div>

        <!-- Main Content -->
        <div class="main-content">
            <div class="posts-header">
                <h2>#general</h2>
            </div>
            <div class="posts-container">
                <div class="post">
                    <div class="post-author">User1</div>
                    <div class="post-content">Hello everyone! 👋</div>
                </div>
                <div class="post">
                    <div class="post-author">User2</div>
                    <div class="post-content">Hey User1! How's it going?</div>
                </div>
                <div class="post">
                    <div class="post-author">User3</div>
                    <div class="post-content">Welcome to the server! 🎉</div>
                </div>
            </div>
            <div class="post-input">
                <input type="text" placeholder="Type your message here..." />
                <button>Send</button>
            </div>
        </div>
    </div>
</body>
</html>`