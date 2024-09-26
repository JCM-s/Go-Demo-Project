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

        const editButtonDiv = document.createElement('div');
        editButtonDiv.setAttribute("class", "editButtonDiv");

        const delete_Blog = document.createElement('button');
        delete_Blog.innerHTML = "Löschen |"
        delete_Blog.setAttribute("class", "editButton");
        delete_Blog.setAttribute("onclick", "deleteBlog("+ jsonData[i].id +")");

        const change_Blog = document.createElement('button');
        change_Blog.innerHTML = "Bearbeiten"
        change_Blog.setAttribute("class", "editButton");
        change_Blog.setAttribute("onclick", "startChangeBlog("+ jsonData[i].id +")");

        editButtonDiv.append(delete_Blog);
        editButtonDiv.append(change_Blog);

        blog.append(blog_title);
        blog.append(blog_autor);
        blog.append(blog_nachricht);
        blog.append(editButtonDiv);

        bloggingArea.append(blog);

        console.log(jsonData[i].id);
    }
}

function addBlog() {
    const autor = prompt("Wie heißt du?");
    const title = prompt("Welchen Titel soll der Blog eintrag haben?");
    const nachricht = document.getElementById('nachricht').value;

    const jsonData = {"title": title, "autor": autor, "nachricht": nachricht};

    if(autor == "" || title == "" || nachricht == "") {
        alert("Bitte fülle alle geforderten Felder aus!");
    } else {
        const xhr = new XMLHttpRequest();
        xhr.open('POST', 'http://localhost:10000/posts', true);
        xhr.setRequestHeader('X-Requested-With', 'XMLHttpRequest');
        xhr.setRequestHeader('Content-Type', 'application/json');
        xhr.onload = function() {
        if (xhr.status === 200) {
            location.reload();
        }
        };
        xhr.send(JSON.stringify(jsonData));
        location.reload();
    }
}

function deleteBlog(id) {
    var result = confirm("Bist du sicher das du den Blogeintrag "+ id +" löschen möchtest?")

    if(result) {
        const xhr = new XMLHttpRequest();
        xhr.open('DELETE', 'http://localhost:10000/posts/' + id + '', true);
        xhr.setRequestHeader('X-Requested-With', 'XMLHttpRequest');
        xhr.s
        xhr.onload = function() {
        if (xhr.status === 200) {
            const jsonData = JSON.parse(xhr.responseText);
            console.log(jsonData);
            location.reload();
        }
        };
        xhr.send(null);
    } else {
        alert("Blog wird nicht gelöscht!");
    }
}

function startChangeBlog(id) {
    const blog = document.getElementById(id);

    
    blog.getElementsByTagName('button')[0].remove();
    blog.getElementsByTagName('button')[0].remove();

    const oldMessage = blog.getElementsByTagName('p')[0].innerHTML;

    const nachricht = document.createElement('textarea');
    nachricht.setAttribute("class", "changeTextarea");
    nachricht.value = oldMessage;
    blog.getElementsByTagName('p')[0].remove();
    
    const saveButton = document.createElement('button');
    saveButton.setAttribute("class", "Button")
    saveButton.setAttribute("onclick", "changeBlog("+ id +")");
    saveButton.innerHTML = "Sichern";

    const cancleButton = document.createElement('button');
    cancleButton.setAttribute("class", "Button");
    cancleButton.setAttribute("onclick", "location.reload()");
    cancleButton.innerHTML = "Abbrechen";

    blog.insertBefore(nachricht, blog.getElementsByTagName('h4')[0].nextSibling);
    blog.append(saveButton);
    blog.append(cancleButton);
}

function changeBlog(id) {
    const blog = document.getElementById(id);

    const title = blog.getElementsByTagName('h2')[0].innerHTML;
    const autor = blog.getElementsByTagName('h4')[0].innerHTML;
    const nachricht = blog.getElementsByTagName('textarea')[0].value;

    const jsonData = {"title": title, "autor": autor, "nachricht": nachricht};

    if(autor == "" || title == "" || nachricht == "") {
        alert("Bitte fülle alle geforderten Felder aus!");
    } else {
        const xhr = new XMLHttpRequest();
        xhr.open('POST', 'http://localhost:10000/posts/' + id, true);
        xhr.setRequestHeader('X-Requested-With', 'XMLHttpRequest');
        xhr.setRequestHeader('Content-Type', 'application/json');
        xhr.onload = function() {
        if (xhr.status === 200) {
            location.reload();
        }
        };
        xhr.send(JSON.stringify(jsonData));
        location.reload();
    }
}

function cancleChangeBlog(id, oldMessage) {
    alert("Blog " + id + " wiht Message: " + oldMessage);
}

main();