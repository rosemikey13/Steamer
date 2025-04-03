const Game = ({game}) => {
    return(
        <div>
            <img src={`https://media.steampowered.com/steamcommunity/public/images/apps/${game.appid}/${game.img_icon_url}.jpg`}/>
            <h1>{game.name}</h1>
            <h2>{game.playtime_forever / 60}</h2>
        </div>
    )
}

export default Game