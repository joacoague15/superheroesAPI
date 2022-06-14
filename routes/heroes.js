import express from "express";

const router = express.Router();

const heroes = [
    {
        id: 1,
        name: 'Batman',
        img: 'https://dam.smashmexico.com.mx/wp-content/uploads/2019/11/batman-1.jpg'
    },
    {
        id: 2,
        name: 'Superman',
        img: 'https://www.cinemascomics.com/wp-content/uploads/2021/06/superman-comics-dest.jpg'
    },
    {
        id: 3,
        name: 'Catwoman',
        img: 'https://i.pinimg.com/originals/5b/a1/46/5ba146b0774396785440c62596f2a2fe.jpg'
    },
    {
        id: 4,
        name: 'Aquaman',
        img: 'https://i.blogs.es/fc993d/aquaman-new-52-comics/1366_2000.jpg'
    },
    {
        id: 5,
        name: 'Iron man',
        img: 'https://sm.ign.com/ign_latam/screenshot/default/ybbpqktez5whedr0-1592031889_31aa.jpg'
    },
    {
        id: 6,
        name: 'Thor',
        img: 'data:image/jpeg;base64,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'
    },
    {
        id: 7,
        name: 'Spiderman',
        img: 'https://uvn-brightspot.s3.amazonaws.com/assets/vixes/n/non_stop_spiderman_marvel_comics.jpg'
    },
    {
        id: 8,
        name: 'Hulk',
        img: 'https://dam.smashmexico.com.mx/wp-content/uploads/2018/11/hulkbio_portada.jpg'
    },
    {
        id: 9,
        name: 'Captain America',
        img: 'https://www.cinemascomics.com/wp-content/uploads/2022/04/capitan-america-comic.jpg'
    },
    {
        id: 10,
        name: 'Flash',
        img: 'https://i.pinimg.com/736x/6d/ee/54/6dee5468061ef77fbf1f374c3b17562c.jpg'
    },
    {
        id: 11,
        name: 'Deadpool',
        img: 'https://sm.ign.com/t/ign_es/gallery/d/deadpools-/deadpools-17-craziest-comic-book-stories_939b.1080.jpg'
    },
    {
        id: 12,
        name: 'Doctor Strange',
        img: 'https://as01.epimg.net/meristation/imagenes/2021/06/07/noticias/1623060203_218082_1623060273_noticia_normal.jpg'
    },
    {
        id: 13,
        name: 'Daredevil',
        img: 'https://www.tonica.la/__export/1607623714044/sites/debate/img/2020/12/10/daredevil-5-comics-conocer-mejor-hombre-sin-miedo_crop1607623664340.jpg_1902800913.jpg'
    },
    {
        id: 14,
        name: 'Wonder Woman',
        img: 'https://dam.smashmexico.com.mx/wp-content/uploads/2017/04/La-historia-de-Wonder-Woman-en-anecdotas-y-curiosidades-cover.jpg'
    },
    {
        id: 15,
        name: 'Black Widow',
        img: 'https://uvn-brightspot.s3.amazonaws.com/assets/vixes/b/black_widow_marvel_comics_3.jpg'
    }
]

const heroesStats = [
    {
        id: 1,
        power: 86,
        durability: 95,
        intelligence: 100
    },
    {
        id: 2,
        power: 97,
        durability: 100,
        intelligence: 70
    },
    {
        id: 3,
        power: 79,
        durability: 60,
        intelligence: 75
    },
    {
        id: 4,
        power: 89,
        durability: 75,
        intelligence: 50
    },
    {
        id: 5,
        power: 84,
        durability: 70,
        intelligence: 100
    },
    {
        id: 6,
        power: 96,
        durability: 90,
        intelligence: 55
    },
    {
        id: 7,
        power: 86,
        durability: 85,
        intelligence: 90
    },
    {
        id: 8,
        power: 85,
        durability: 80,
        intelligence: 30,
    },
    {
        id: 9,
        power: 85,
        durability: 92,
        intelligence: 60
    },
    {
        id: 10,
        power: 85,
        durability: 50,
        intelligence: 85
    },
    {
        id: 11,
        power: 82,
        durability: 95,
        intelligence: 60
    },
    {
        id: 12,
        power: 88,
        durability: 35,
        intelligence: 85
    },
    {
        id: 13,
        power: 83,
        durability: 90,
        intelligence: 80
    },
    {
        id: 14,
        power: 95,
        durability: 95,
        intelligence: 70
    },
    {
        id: 15,
        power: 76,
        durability: 65,
        intelligence: 75
    }
]

router.get('/', (req, res) => {
    res.send(heroes);
});

router.get('/:id/stats', (req, res) => {
    const statID = parseInt(req.params.id)


    for (let heroStat of heroesStats) {
        if (heroStat.id === statID)
            res.send(heroStat);
    }
});

export default router;