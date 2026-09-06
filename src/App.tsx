import Tasks from "./Components/Tasks";
import AddTask from "./Components/AddTask";
import { useEffect, useState } from "react";
import { v4 } from "uuid";

function App() {
    const [tasks, setTasks] = useState(
        JSON.parse(localStorage.getItem("tasks")) || []
    );

  
    useEffect(() => {
        async function fetchTasks() {
            const response = await fetch(
                "https://jsonplaceholder.typicode.com/todos?_limit=10"
            );

            const data = await response.json();

            const apiTasks = data.map((task) => ({
                id: task.id,
                title: task.title,
                description: "Tarefa carregada da API",
                isCompleted: false, 
            }));

            setTasks(apiTasks);
        }
        const savetask = localStorage.getItem("tasks");
        if (!savetask) {

        fetchTasks();
        }
    }, []);

   
    useEffect(() => {
        localStorage.setItem("tasks", JSON.stringify(tasks));
    }, [tasks]);

   
    function OnTaskClick(TaskId) {
        const newTasks = tasks.map((task) => {
            if (task.id === TaskId) {
                return {
                    ...task,
                    isCompleted: !task.isCompleted,
                };
            }

            return task;
        });

        setTasks(newTasks);
    }

    
    function OnDeleteTaskClick(TaskId) {
        const newTasks = tasks.filter(
            (task) => task.id !== TaskId
        );

        setTasks(newTasks);
    }

    
    function OnAddTaskClick({ title, description }) {
        const newTask = {
            id: v4(),
            title: title,
            description: description,
            isCompleted: false,
        };

        setTasks([...tasks, newTask]);
    }

    return (
        <div className="w-screen min-h-screen bg-slate-500 flex justify-center p-6">
            <div className="w-[500px] space-y-4">
                <h1 className="text-3xl text-slate-100 font-bold text-center">
                    Gerenciador de tarefas
                </h1>

                <AddTask
                    OnAddTaskClick={OnAddTaskClick}
                />

                <Tasks
                    tasks={tasks}
                    onTaskClick={OnTaskClick}
                    onDeleteClick={OnDeleteTaskClick}
                />
            </div>
        </div>
    );
}

export default App;