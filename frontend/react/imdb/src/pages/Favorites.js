import React, {useContext} from 'react';
import FavoriteList from '../components/movies/FavoriteList';
import {gql, useQuery} from "@apollo/client";

// this function component get user's favorite list data and send it to another component that shows those movies on by one
function FavoritesPage(props) {
    const GET_FAVORITES = gql`
        query FavoritesOfUser($userID : ID!){
            favoritesOfUser(userID : $userID){
                movieID
                movieTitle
                movieImage
            }
        }
    `;

    let userID = props.userID

    const {loading, error, data} = useQuery(GET_FAVORITES,
        {
            variables: {
                userID: userID || 0
            }
        })

    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error :</p>;

    let content
    if (data["favoritesOfUser"] === null) {
        content = <p
            style={{color: "yellow"}}>You got no favorites yet. Start adding some?</p>;
    } else {
        content = <FavoriteList movies={data}/>;
    }

    return (
        <section>
            <h1 style={{color: "yellow"}}>My Favorites</h1>
            {content}
        </section>
    );
}

export default FavoritesPage;