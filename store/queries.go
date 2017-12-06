package store

const GetSpotByID = `SELECT * FROM parkingSpot WHERE id = ?`
const GetAllSpotsWithLimit = `SELECT * FROM parkingSpot LIMIT 10`
const GetAllSpotsNoLimit = `SELECT * FROM parkingSpot`

const EarthRadius = 3959
