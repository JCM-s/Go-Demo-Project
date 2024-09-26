function test() {
    alert("Hallo Welt!");
}

function main() {
    const xhr = new XMLHttpRequest();
    xhr.open('GET', 'http://localhost:10000/posts', true);
    xhr.setRequestHeader('X-Requested-With', 'XMLHttpRequest');
    xhr.onload = function() {
    if (xhr.status === 200) {
        const jsonData = JSON.parse(xhr.responseText);
        console.log(jsonData);
        showBlogs(jsonData);
    }
    };
    xhr.send();
}

function showBlogs(jsonData) {
    const bloggingArea = document.getElementById('BloggingArea');

    for (var i = 0; i < jsonData.length; i++) {

        const blog = document.createElement('div');
        blog.setAttribute("class", "Blog");
        blog.setAttribute("id", jsonData[i].id);

        const blog_title = document.createElement('h2');
        blog_title.setAttribute("class", "blogTitle");
        blog_title.innerHTML = jsonData[i].title;

        const blog_autor = document.createElement('h4');
        blog_autor.setAttribute("class", "blogAutor");
        blog_autor.innerHTML = jsonData[i].autor;

        const blog_nachricht = document.createElement('p');
        blog_nachricht.setAttribute("class", "blogNachricht")
        blog_nachricht.innerHTML = jsonData[i].nachricht;

        blog.append(blog_title);
        blog.append(blog_autor);
        blog.append(blog_nachricht);

        bloggingArea.append(blog);

        console.log(jsonData[i].id);
    }
}

function addBlog() {
    const autor = prompt("Was ist dein Name");
    const title = prompt("Welchen Titel soll der Blog eintrag haben?");
    const nachricht = document.getElementById('nachricht').value;

    console.log(nachricht);

    const jsonData = {"title": title, "autor": autor, "nachricht": nachricht};

    if(autor == "" || title == "" || nachricht == "") {
        alert("Bitte fülle alle geforderten Felder aus!");
    } else {
        const xhr = new XMLHttpRequest();
        xhr.open('POST', 'http://localhost:10000/posts', true);
        xhr.setRequestHeader('X-Requested-With', 'XMLHttpRequest');
        xhr.setRequestHeader('Content-Type', 'application/json');
        xhr.s
        xhr.onload = function() {
        if (xhr.status === 200) {
            const jsonData = JSON.parse(xhr.responseText);
            console.log(jsonData);
            showBlogs(jsonData);
        }
        };
        xhr.send(JSON.stringify(jsonData));
        location.reload();
    }
}

main();