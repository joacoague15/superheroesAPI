import express from 'express';
import bodyParser from 'body-parser';

import heroesRoutes from './routes/heroes.js'

const app = express();
const PORT = 4000;

app.use(bodyParser.json());

app.get('/', (req, res) => {
    console.log('TEST')
    res.send('HI')
});

app.use('/heroes', heroesRoutes);

app.listen(PORT, () => {
    console.log(`Server running on port http://localhost:${PORT}`)
})