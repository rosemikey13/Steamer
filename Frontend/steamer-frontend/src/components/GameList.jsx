import Game from "./Game"

const GameList = ({games}) => {
    return (
        games.Map(game => <Game game={game}/>)
    )
}

export default GameList