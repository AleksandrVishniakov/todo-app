export default class TaskModel {
  constructor() {
    this.ended = false
  }

  setParams(description, dueDate, color) {
    this.ended = true

    this.description = description
    this.dueDate = dueDate,
    this.color = color
  }
}

const getRandomInteger = (a = 0, b = 1) => {
    const lower = Math.ceil(Math.min(a, b));
    const upper = Math.floor(Math.max(a, b));
  
    return Math.floor(lower + Math.random() * (upper - lower + 1));
};

const generateDescription = () => {
  const descriptions = [
    'Погулять с собакой',
    'Сделать домашку',
    'Приготовить подарок маме',
  ];

  const randomIndex = getRandomInteger(0, descriptions.length - 1);

  return descriptions[randomIndex];
};

export const generateTask = () => {
  const newTask = new TaskModel()
  newTask.setParams(generateDescription(), '20.02.10', 'black')
  return newTask
};