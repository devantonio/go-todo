const _$ = document.querySelector.bind(document);
const _$$ = document.querySelectorAll.bind(document);

document.addEventListener("DOMContentLoaded", function() {
    console.log('DOM fully loaded and parsed');

    let createTaskForm = document.getElementById('create-task-form');

    createTaskForm.onsubmit = (e) => {
        e.preventDefault();

        // Get the value of the task title input
        const titleInput = document.getElementById('task-title');
        const taskTitle = titleInput.value;

        const descriptionInput = document.getElementById('task-description');
        const taskDescription = descriptionInput.value;

        const completedInput = document.getElementById('task-completed');
        const taskCompleted = completedInput.checked;

        // Check if the task title is not empty
        if (taskTitle.trim() === '') {
            alert('Please enter a task title.');
            return;
        }

        // Send the task data to the server
        fetch('/tasks', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ title: taskTitle, description: taskDescription, completed: taskCompleted}),
        })
        .then(response => response.json())
        .then(data => {
            console.log('Task added:', data);
            let taskItem = document.createElement("li");
                taskItem.innerText = `${data.Title}: ${data.Description}`;

            // append the task to the list of tasks
           _$('.template-tasks > ul').appendChild(taskItem);


            // Optionally, you can clear the input field and update the UI
            titleInput.value = '';
            descriptionInput.value = '';
            completedInput.value = '';

        })
        .catch(error => console.error('Error adding task:', error));
    };

    // fetch('/tasks')
    //     .then(response => {
    //         console.log('Fetch response:', response);
    //         return response.json();
    //     })
    //     .then(data => {
    //         console.log('Fetched data:', data);

    //         const container = document.querySelector('.task-container');
    //         if (container) {
    //             data.forEach(task => {
    //                 console.log('Task:', task);
    //                 const taskElement = document.createElement('div');
    //                 taskElement.innerText = task.Title;
    //                 container.appendChild(taskElement);
    //             });
    //         } else {
    //             console.error('Container element not found');
    //         }
    //     })
    //     .catch(error => console.error('Error fetching tasks:', error));
});
