const form = document.querySelector("#new-task-form");
const input = document.querySelector("#new-task-input");
const list_el = document.querySelector("#tasks");

loadWebsite()


// Adds an event listener to execute tasks when the website is loaded
function loadWebsite(){
    window.addEventListener("load", () =>{

        loadUserTasks();

        form.addEventListener("submit", (e) => {
            e.preventDefault();

            const task = input.value;

            if(!task){
                alert("Gib bitte eine Aufgabe an!");
                return;
            }

            addTask(-1, task)
        });
    });
}

//Adds a task element into the list
function addTask(id=-1, taskTitle="", description="", isDone=false){
    const task_el = document.createElement("div");
    task_el.classList.add("task");
    task_el.setAttribute("data-id", id);

    const task_content_el = document.createElement("div");
    task_content_el.classList.add("content");

    task_el.appendChild(task_content_el);

    const task_checker_el = document.createElement("img");
    if(isDone) {
        task_checker_el.src = "/img/checked.png";
    }else{
        task_checker_el.src = "/img/unchecked.png";
    }
    task_content_el.appendChild(task_checker_el);

    const task_input_el = document.createElement("input")
    task_input_el.classList.add("text");
    task_input_el.type = "text";
    task_input_el.value = taskTitle;
    task_input_el.setAttribute("readonly", "readonly")

    task_content_el.appendChild(task_input_el)

    const task_actions_el = document.createElement("div");
    task_actions_el.classList.add("actions");

    const task_edit_el = document.createElement("button");
    task_edit_el.classList.add("edit");
    task_edit_el.innerHTML = "Bearbeiten";

    const task_delete_el = document.createElement("button");
    task_delete_el.classList.add("delete");
    task_delete_el.innerHTML = "Löschen";

    task_actions_el.appendChild(task_edit_el);
    task_actions_el.appendChild(task_delete_el);

    task_el.appendChild(task_actions_el);

    list_el.appendChild(task_el);

    const task_description_el = document.createElement("div");
    task_description_el.classList.add("description");

    const description_input_el = document.createElement("textarea");
    description_input_el.classList.add("text");
    description_input_el.value = description;
    description_input_el.placeholder = "Beschreibung hinzufügen...";
    description_input_el.setAttribute("readonly", "readonly");

    task_description_el.appendChild(description_input_el);
    task_el.appendChild(task_description_el);

    input.value = "";

    toggleEdit(task_edit_el, task_input_el, description_input_el);
    deleteTask(task_delete_el, task_el);
    toggleCheckbox(task_checker_el);
}

//Adds an event listener to toggle between edit and static display
function toggleEdit(task_edit_el, task_input_el, description_input_el){
    task_edit_el.addEventListener("click", () => {
        if(task_edit_el.innerText.toLocaleLowerCase() === "bearbeiten"){
            task_input_el.removeAttribute("readonly");
            description_input_el.removeAttribute("readonly");
            task_input_el.focus();
            task_edit_el.innerText = "Speichern";
        }else{
            task_input_el.setAttribute("readonly", "readonly");
            description_input_el.setAttribute("readonly", "readonly");
            task_edit_el.innerText = "Bearbeiten";
        }
    })
}

//Adds an event listener to delete the targeted task element
function deleteTask(task_delete_el, task_el){
    task_delete_el.addEventListener("click", () => {

        const URL = "/todo-list/user1/tasks/" + task_el.getAttribute("data-id");

        fetch(URL, {method: "DELETE"})
            .then(response => {
                if(!response.ok){
                    //Nachricht einfügen
                    throw new Error();
                }
                list_el.removeChild(task_el);
            })
            .catch(error => {
                console.log(error);
            })
    });
}

//Adds an event listener to toggle and mark the checkbox of tasks
function toggleCheckbox(task_checker_el){
    task_checker_el.addEventListener("click", function (e) {
        // if(e.target.tagName === "IMG"){
        //     e.target.classList.toggle("checked");
        // }
        if(task_checker_el.getAttribute("src") === "/img/checked.png"){
            task_checker_el.setAttribute("src", "/img/unchecked.png");
        }else{
            task_checker_el.setAttribute("src", "/img/checked.png");
        }
    })
}

//Fetches and loads user specific tasks
function loadUserTasks(){
    const URL = "http://localhost:8080/todo-list/user1/tasks";

    fetch(URL)
        .then(response => {
            if(!response.ok){
                //Nachricht überarbeiten
                throw new Error("Network response was not ok");
            }
            return response.json();
        })
        .then(data => {
            data.forEach(task => {
                addTask(task.id, task.title, task.description, task.is_done);
            })
        })
        .catch(error => {
            console.error("Fetch error:", error);
        })
}