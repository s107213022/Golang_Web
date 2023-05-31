<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <!-- 引入 Bootstrap 的 CSS 檔案 -->
    <link rel="stylesheet" href="/static/dist/css/bootstrap.min.css">
    <title>Bootstrap Frontend</title>
</head>
<body>
    <div class="container">
        <h1>Welcome to My Website</h1>
        <p>This is a sample frontend built with Bootstrap.</p>
        
        <div class="row">
            <div class="col-md-6">
                <h1>About</h1>
                <p>Hello, {{.Name}}</p> 
            </div>
            <div class="col-md-6">
                <h2>Contact Us</h2>
                <form>
                    <div class="mb-3">
                        <label for="name" class="form-label">Name</label>
                        <input type="text" class="form-control" id="name" placeholder="Enter your name">
                    </div>
                    <div class="mb-3">
                        <label for="email" class="form-label">Email</label>
                        <input type="email" class="form-control" id="email" placeholder="Enter your email">
                    </div>
                    <div class="mb-3">
                        <label for="message" class="form-label">Message</label>
                        <textarea class="form-control" id="message" rows="3" placeholder="Enter your message"></textarea>
                    </div>
                    <button type="submit" class="btn btn-primary">Submit</button>
                </form>
            </div>
        </div>
    </div>

    <!-- 引入 Bootstrap 的 JavaScript 檔案（需要包含在 body 結束標籤之前） -->
    <script src="/static/dist/js/bootstrap.bundle.min.js"></script>
</body>
</html>