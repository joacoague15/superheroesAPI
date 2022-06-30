import express from 'express';
import bodyParser from 'body-parser';
import heroesRoutes from './routes/heroes.js';
import { db } from './config/database.js';

try {
    await db.authenticate();
    console.log('Connection has been established successfully.');
} catch (error) {
    console.error('Unable to connect to the database:', error);
}

const app = express();
const PORT = 4000;

app.use(bodyParser.json());

app.use('/heroes', heroesRoutes);

app.listen(PORT, () => {
    console.log(`Server running on port http://localhost:${PORT}`)
})