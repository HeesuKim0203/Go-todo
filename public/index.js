const button = document.querySelector('.todo-list-add-btn') ;
const todoListItem = document.querySelector('.todo-list') ;
const todoListInput = document.querySelector('.todo-list-input') ;

button.addEventListener('click', (event) => {
    event.preventDefault();
    
    const item = document.querySelector('.todo-list-input').value ;

    if(item) {
        const formdata = new FormData();
        formdata.append("name", item);

        fetch("/todos", {
            method : 'POST',
            redirect : 'follow',
            body : formdata
        })
        .then(response => response.text())
        .then(result => {
            const todoData = JSON.parse(result) ;
            addItem(todoData.name, todoData.completed, todoData.id) ;
        }).catch(error => console.log('error', error)) ;

        Array.prototype.forEach.call(document.getElementsByTagName('li'), (element) => {
            if( !element.onclick ) {
                element.onclick = todoListLiOnClickEvent ;
            }
        }) ;

        todoListInput.value = "" ;
    }
}) ;

function todoListLiOnClickEvent(event) {

    const target = event.currentTarget ;
    const targetCheckElement = target.querySelector('input') ;
    let complete = true ;

    if( targetCheckElement.checked ) complete = false ;

    fetch(`/complete-todo/${target.id}?complete=${complete}`, {
        method : 'GET',
        redirect : 'follow',
    })
    .then(response => response.text())
    .then(result => {
        const { success } = JSON.parse(result) ;
        if( success ) {
            if( complete ) {
                targetCheckElement.setAttribute('checked', 'checked') ;
                target.className += 'completed' ;
            }else {
                targetCheckElement.removeAttribute('checked') ;
                target.className = target.className.replace('completed', '') ;
            }
        }
    })
    .catch(error => console.log('error', error)) ;
}

todoListItem.addEventListener('click', (event) => {
    if( event.target.className.includes('remove') ) {
        const id = event.target.parentNode.id ;

        if( !id ) return ;

        fetch(`/todos/${event.target.parentNode.id}`, {
            method : 'DELETE',
            redirect : 'follow',
        })
        .then(response => response.text())
        .then(result => {
            const { success } = JSON.parse(result) ;
            if( success ) todoListItem.removeChild(event.target.parentNode) ;
        })
        .catch(error => console.log('error', error)) ;
    }
}, true) ;

function addItem(name, completed, id) {
    todoListItem.innerHTML += `<li ${ completed ? "class = 'completed'" : ""} ${ id ? "id = " + id : ""}><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' ${ completed ? "checked = 'checked'" : ""} />${ name }<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>` ;
}

(function init() {
    fetch("http://localhost:3001/todos", {
        method : 'GET',
        redirect : 'follow'
    })
    .then(response => response.text())
    .then(result => {
        const todoListData = JSON.parse(result) ;
        todoListData.forEach((todo) => {
            if(todo) {
                addItem(todo.name, todo.completed, todo.id) ;
            }
        }) ; 

        Array.prototype.forEach.call(document.getElementsByTagName('li'), (element) => {
            if( !element.onclick ) {
                element.onclick = todoListLiOnClickEvent ;
            }
        })
    })
    .catch(error => console.log('error', error)) ;
})() ;