const form = document.querySelector("#new-task-form");
const input = document.querySelector("#new-task-input");
const list_el = document.querySelector("#tasks");

loadWebsite()

// Adds an event listener to execute tasks when the website is loaded
function loadWebsite(){
    window.addEventListener("load", () =>{
        loadUserTasks();
        submitTask();
    });
}

//Adds a task element with functionality and event listener into the list
function addTask(id="", taskTitle="", description="", isDone=false){

    //Container to place components of task elements
    const task_el = document.createElement("div");
    task_el.classList.add("task");
    task_el.setAttribute("data-id", id);

    const task_content_el = document.createElement("div");
    task_content_el.classList.add("content");

    task_el.appendChild(task_content_el);

    //Checkbox of task element
    const task_checker_el = document.createElement("img");
    if(isDone) {
        task_checker_el.src = "/img/checked.png";
        task_checker_el.setAttribute("checked", "true");
    }else{
        task_checker_el.src = "/img/unchecked.png";
        task_checker_el.setAttribute("checked", "false");
    }

    task_content_el.appendChild(task_checker_el);

    //Title of task element
    const task_input_el = document.createElement("input")
    task_input_el.classList.add("text");
    task_input_el.type = "text";
    task_input_el.value = taskTitle;
    task_input_el.setAttribute("readonly", "readonly")

    task_content_el.appendChild(task_input_el)

    //Action buttons to modify or delete task elements
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

    //Description section of task element
    const task_description_el = document.createElement("div");
    task_description_el.classList.add("description");

    const description_input_el = document.createElement("textarea");
    description_input_el.classList.add("text");
    description_input_el.value = description;
    description_input_el.placeholder = "Beschreibung hinzufügen...";
    description_input_el.setAttribute("readonly", "readonly");

    task_description_el.appendChild(description_input_el);
    task_el.appendChild(task_description_el);

    //Category section of task element
    const categories_el = document.createElement("div");
    categories_el.classList.add("categories");

    const categories_title_container_el = document.createElement("div");
    categories_title_container_el.classList.add("title");

    const categories_image_el = document.createElement("img");
    categories_image_el.src = "/img/tag.png";

    const categories_title_el = document.createElement("h2");
    categories_title_el.innerHTML = "Kategorien";

    categories_title_container_el.appendChild(categories_image_el);
    categories_title_container_el.appendChild(categories_title_el);

    const categories_content_el = document.createElement("div");
    categories_content_el.classList.add("category-content");

    const categories_list_el = document.createElement("ul");
    const categories_input_el = document.createElement("input");
    categories_input_el.classList.add("text");
    categories_input_el.setAttribute("readonly", "readonly");

    categories_list_el.appendChild(categories_input_el);
    categories_content_el.append(categories_list_el);
    categories_el.appendChild(categories_title_container_el);
    categories_el.appendChild(categories_content_el);

    task_el.appendChild(categories_el);

    input.value = "";

    addTags(categories_list_el, categories_input_el);
    toggleEdit(id, task_edit_el, task_input_el, description_input_el, categories_input_el, task_checker_el.getAttribute("checked"));
    deleteTask(task_delete_el, task_el);
    toggleCheckbox(id, task_input_el.value, description_input_el.value, task_checker_el);
}

//Fetches and loads user specific tasks from the database
function loadUserTasks(){
    const URL = "http://localhost:8080/user1/tasks";

    fetch(URL)
        .then(response => {
            if(!response.ok){
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

//Adds an event listener to submit new empty tasks and create a new set of data in the database
function submitTask(){
    form.addEventListener("submit", (e) => {
        e.preventDefault();

        const task = input.value;

        if(!task){
            alert("Gib bitte eine Aufgabe an!");
            return;
        }

        const URL = "/user1/tasks";

        const data = {
            title: task,
            description: "",
            is_done: false
        }

        fetch(URL, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(data)
        })
            .then(response => {
                if(!response.ok){
                    throw new Error("Network response was not ok");
                }
                addTask("", task)
                return response.json();
            })
            .then(data => {
                console.log("Response successful:", data);
            })
            .catch(error => {
                console.error(error);
            })
    });
}

//Updates a task and makes changes in the database
function updateTask(id, title, description, is_done){
    const URL = "/todo-list/user1/tasks/" + id;

    const data = {
        title: title,
        description: description,
        is_done: {"true": true, "false": false}[is_done]
    }

    fetch(URL, {
        method: "PATCH",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
    })
        .then(response => {
            if(!response.ok){
                throw new Error("Network response was not ok");
            }
            return response.json();
        })
        .then(data => {
            console.log("Response successful:", data);
        })
        .catch(error => {
            console.error(error);
        })
}

//Returns all current tags in the category section
function getTags(categories_list_el){
    let tags = [];
    categories_list_el.querySelectorAll("li").forEach(li => tags.push(li.innerText));
    return tags;
}

//Adds an event handler to confirm and insert tags into the category section
function addTags(categories_list_el, categories_input_el){
    categories_input_el.addEventListener("keyup",(e) => {
        if(e.key == "Enter"){
            let tags = getTags(categories_list_el);
            let tag = e.target.value.replace(/\s+/g, " ");
            if(tag.length > 1 && !tags.includes(tag)){
                tag.split(",").forEach(tag => {
                    tags.push(tag);
                    createTag(categories_list_el, tags);
                })
            }
            e.target.value = "";
        }
    })
}

//Creates a tag and inserts it into the right position
function createTag(categories_list_el, tags){
    categories_list_el.querySelectorAll("li").forEach(li => li.remove());
    tags.slice().reverse().forEach(tag => {
        let liTag = `<li>${tag}<i class="uit uit-multiply" onclick="removeTag(this)"></i></li>`;
        categories_list_el.insertAdjacentHTML("afterbegin", liTag);
    })
}

//Handler for onclick event to delete a tag
function removeTag(element){
    element.parentElement.remove();
}

//Adds an event listener to toggle between edit and static display. Changes update the database
function toggleEdit(id, task_edit_el, task_input_el, description_input_el, categories_input_el, is_done){
    task_edit_el.addEventListener("click", () => {
        if(task_edit_el.innerText.toLocaleLowerCase() === "bearbeiten"){
            task_input_el.removeAttribute("readonly");
            description_input_el.removeAttribute("readonly");
            categories_input_el.removeAttribute("readonly");
            task_input_el.focus();
            task_edit_el.innerText = "Speichern";
        }else{
            task_input_el.setAttribute("readonly", "readonly");
            description_input_el.setAttribute("readonly", "readonly");
            categories_input_el.setAttribute("readonly", "readonly");
            task_edit_el.innerText = "Bearbeiten";
            updateTask(id, task_input_el.value, description_input_el.value, is_done);
        }
    })
}

//Adds an event listener to delete the targeted task element and make changes in the database
function deleteTask(task_delete_el, task_el){
    task_delete_el.addEventListener("click", () => {

        const URL = "/todo-list/user1/tasks/" + task_el.getAttribute("data-id");

        fetch(URL, {method: "DELETE"})
            .then(response => {
                if(!response.ok){
                    throw new Error("Network response was not ok");
                }
                list_el.removeChild(task_el);
            })
            .catch(error => {
                console.log(error);
            })
    });
}

//Adds an event listener to toggle and mark the checkbox of tasks + update in database
function toggleCheckbox(id, title, description, task_checker_el){
    task_checker_el.addEventListener("click", () => {
        if(task_checker_el.getAttribute("src") === "/img/checked.png"){
            task_checker_el.setAttribute("src", "/img/unchecked.png");
            task_checker_el.setAttribute("checked", "false");
        }else{
            task_checker_el.setAttribute("src", "/img/checked.png");
            task_checker_el.setAttribute("checked", "true");
        }
        updateTask(id, title, description, task_checker_el.getAttribute("checked"))
    })
}
