import express from "express";

const router = express.Router();

const heroes = [
    {
        name: 'Batman'
    },
    {
        name: 'Superman'
    },
    {
        name: 'Catwoman'
    },
    {
        name: 'Aquaman'
    },
    {
        name: 'Iron man'
    },
    {
        name: 'Thor'
    },
    {
        name: 'Spiderman'
    },
    {
        name: 'Hulk'
    },
    {
        name: 'Captain America'
    },
    {
        name: 'Flash'
    }
]

router.get('/', (req, res) => {
    res.send(heroes);
});

export default router;