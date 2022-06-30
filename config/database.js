import { Sequelize } from "sequelize";

export const db = new Sequelize('postgres://joaquinaguero@localhost:5432/superheroesDB')