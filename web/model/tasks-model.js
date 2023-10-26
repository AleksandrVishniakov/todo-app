import {generateTask} from './task.js';

export default class TasksModel {
  tasks = Array.from({length: 4}, generateTask);

  getTasks = () => this.tasks;

  addTask = (newTask) => {
    this.tasks.push(newTask)
  }

  removeTask = (task) => {
    this.tasks.splice(this.tasks.indexOf(task), 1)
    console.log(this.tasks)
  }
}