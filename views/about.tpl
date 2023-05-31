<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Black Theme</title>
    <!-- 引入 Bootstrap 的 CSS 檔案 -->
    <link rel="stylesheet" href="/static/dist/css/bootstrap.min.css">
    <style>
        /* 自訂的黑色主題樣式 */
        body {
            background-color: #000;
            color: #fff;
        }
    </style>
</head>
<body>
    <h1>About</h1>
    <p>Hello, {{.Name}}</p>  

    <!-- 引入 Bootstrap 的 JavaScript 檔案（需要包含在 body 結束標籤之前） -->
    <script src="/static/dist/js/bootstrap.bundle.min.js"></script>
</body>
</html>