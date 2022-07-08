import express from "express";

const router = express.Router();

import { SuperheroesRecruitment } from "../models/SuperheroesRecruitment.js";

router.get('/', (req, res) => {
    const { page, pageSize } = req.query

    const offset = (page - 1) * pageSize
    const limit = pageSize



    SuperheroesRecruitment.findAll({
        limit,
        offset
    })
        .then(heroes => {
            const heroesQueried = []
            heroes.map(hero => {
                heroesQueried.push({
                    id: hero.id,
                    name: hero.c_name,
                    power: hero.c_power,
                    durability: hero.c_durability,
                    intelligence: hero.c_intelligence,
                    img: hero.c_imgURL
                })
            })
            console.log('HEROES QUERIED: ', heroesQueried)
            res.send(heroesQueried);
        })
        .catch(err => console.log(err))
});

router.get('/:id/stats', (req, res) => {
    const statID = parseInt(req.params.id)

    SuperheroesRecruitment.findOne({ where: { id: statID }  })
        .then(heroStat => res.send({
            id: heroStat.id,
            name: heroStat.c_name,
            power: heroStat.c_power,
            durability: heroStat.c_durability,
            intelligence: heroStat.c_intelligence,
        }))
        .catch(err => console.log(err))

});

router.post('/create', (req, res) => {
    const createdHero = req.body

    const id = createdHero.id
    const name = createdHero.name
    const power = parseInt(createdHero.power)
    const durability = parseInt(createdHero.durability)
    const intelligence = parseInt(createdHero.intelligence)
    const imgURL = createdHero.imgURL

    console.log('IMG URL: ', createdHero.imgURL)

    SuperheroesRecruitment.create({
        c_id: id,
        c_name: name,
        c_power: power,
        c_durability: durability,
        c_intelligence: intelligence,
        c_imgURL: imgURL
    })
        .then(() => console.log('New hero added to DB'))
        .catch(err => console.log(err))

    res.status(201).send(req.body)
});

export default router;