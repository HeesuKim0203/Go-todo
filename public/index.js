const button = document.querySelector('.todo-list-add-btn') ;
const todoListItem = document.querySelector('.todo-list') ;
const todoListInput = document.querySelector('.todo-list-input') ;

button.addEventListener('click', (event) => {
    event.preventDefault();
    
    const item = document.querySelector('.todo-list-input').value ;

    if (item) {
        todoListItem.innerHTML += ("<li><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' />" + item + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
        todoListInput.value = "" ;
    }
}) ;

todoListItem.addEventListener('change', (event) => {
    if( event.target.checked !== undefined ) {
        if( event.target.getAttribute('checked') === 'checked' ) {
            event.target.removeAttribute('checked') ;
            event.target.parentNode.className = event.target.parentNode.className.replace(' completed', '') ;
        } else {
            event.target.setAttribute('checked', 'checked') ;
            event.target.parentNode.className += ' completed' ;
        }
    }
    console.log(event.target) ;
}) ;

todoListItem.addEventListener('click', (event) => {
    if( event.target.className.includes('remove') ) {
        todoListItem.removeChild(event.target.parentNode) ;
    }
}) ;