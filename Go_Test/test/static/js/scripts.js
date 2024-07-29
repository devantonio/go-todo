// Example JavaScript
document.addEventListener('DOMContentLoaded', function() {
    fetch('/tasks')
        .then(response => response.json())
        .then(data => {
            const container = document.querySelector('.container');
            data.forEach(task => {
                const taskElement = document.createElement('div');
                taskElement.innerText = task.title;
                container.appendChild(taskElement);
            });
        })
        .catch(error => console.error('Error fetching tasks:', error));
});
