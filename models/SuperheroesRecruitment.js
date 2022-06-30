import { DataTypes } from "sequelize";
import { db } from '../config/database.js';

export const SuperheroesRecruitment = db.define('t_heroes_to_recruit', {
    c_name: {
        type: DataTypes.STRING,
        allowNull: false
    },
    c_power: {
        type: DataTypes.INTEGER,
        allowNull: false
    },
    c_durability: {
        type: DataTypes.INTEGER,
        allowNull: false
    },
    c_intelligence: {
        type: DataTypes.INTEGER,
        allowNull: false
    },
    c_img: {
        type: DataTypes.STRING.BINARY,
        allowNull: false
    }
}, {
    updatedAt: false
})

SuperheroesRecruitment.sync()
    .then(() => console.log('GREAT!'))
    .catch(err => console.log(err))