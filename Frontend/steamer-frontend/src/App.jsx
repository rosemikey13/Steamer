import { useEffect, useState } from 'react'
import './App.css'
import GameList from './components/GameList'

function App() {
  const [games, setGames] = useState([])

  useEffect(() => {
    const getGames = async() => {
      try {
        const req = await fetch("http://localhost:8069")
        const res = await req.json()
        setGames([...res.games])
        
      } catch (error) {
        console.log(error)
      }
    }
    getGames()
  }, [])

  return (
    <>
    {games.length > 0 ? <GameList games={games}/> : <h1>No Games Yet!</h1>}
    </>
  )
}

export default App
