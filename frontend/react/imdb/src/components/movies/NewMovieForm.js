import {useState} from "react";

import Card from "../ui/Card";
import classes from "./NewMovieForm.module.css";

import {styled} from "@mui/material/styles";
import {gql, useMutation, useQuery} from "@apollo/client";
import {Stack} from "@mui/material";
import Button from "@mui/material/Button";

function NewMovieForm() {
    let NEW

    let NEW_MOVIE = gql`
        mutation CreateMovie ($title: String!, $description: String!, $genre: String!, $rank: Int!, $director_id: ID!, $image: String!, $topic: String!, $text: String!) {
            createMovie(movie: {title: $title , description: $description, genre: $genre , rank: $rank, director_id: $director_id, image: $image, topic: $topic, text: $text}) {
                id
            }
        }
    `;

    const DIRECTOR_ID = gql`
        query DirectorIdByName($director: String!) {
            directorIdByName(name: $director)
        }
    `;

    let NEW_MOVIE_AND_DIRECTOR = gql`
        mutation CreateMovieAndDirector ($title: String!, $description: String!, $genre: String!, $rank: Int!, $director_name: String!, $image: String!, $topic: String!, $text: String! $profileImage: String!, $bornAt: String!){
            createMovieAndDirector(title: $title , description: $description, genre: $genre , rank: $rank, directorName: $director_name, image: $image, topic: $topic, text: $text, profileImage: $profileImage, bornAt: $bornAt) {
                id
            }
        }
    `;


    const [givenTitle, setTitle] = useState('')
    const [givenDirector, setDirector] = useState('')
    const [givenDescription, setDescription] = useState('')
    const [givenText, setText] = useState('')
    const [givenRank, setRank] = useState('')
    const [givenWorth, setWorth] = useState('')
    const [givenGenre, setGenre] = useState('')
    const [givenImage1, setImage1] = useState('')
    const [givenImage2, setImage2] = useState('')
    const [givenTopic, setTopic] = useState('')


    const Input = styled("input")({
        display: "none",
    });

    const id = useQuery(DIRECTOR_ID,
        {
            variables: {
                director: givenDirector,
            }
        }).data

    let int = NaN
    let index = NaN
    let data = JSON.stringify(id)
    let exist
    let unique

    if (id && id["directorIdByName"]) {
        index = data.indexOf(":")
        int = parseInt(data.slice(index + 2, data.length - 2), 10)
        console.log(int)
        NEW = NEW_MOVIE
        exist = true
    } else {
        NEW = NEW_MOVIE_AND_DIRECTOR
        exist = false
    }

    const [addMovie] = useMutation(NEW,
        {
            variables: {
                title: givenTitle,
                image: givenImage1 || givenImage2 ||'https://pharem-project.eu/wp-content/themes/consultix/images/no-image-found-360x250.png',
                profileImage: "https://hope.be/wp-content/uploads/2015/05/no-user-image.gif",
                bornAt: "1.1.1111",
                description: givenDescription,
                review: givenText,
                rank: givenRank,
                worth: givenWorth,
                genre: givenGenre,
                topic :givenTopic,
                text: givenText,
                director_id: int,
                director_name: givenDirector,
            },
            onCompleted: function (data) {
                console.log("data:", data)
                if (exist === true) {
                    unique = data["createMovie"]["id"]
                } else {
                    unique = data["createMovieAndDirector"]["id"]
                }
                return window.location.replace("/moviePage/" + unique)

            },
            onError: function (error) {
                console.log("error:",error)
            }
        });


    return (
        <Card>
            <form className={classes.form} >
                <div className={classes.control}>
                    <label htmlFor="title">Movie Title</label>
                    <input type="text" datatype="String" required id="title" value={givenTitle} onChange={event => setTitle(event.target.value)}/>
                </div>

                <div className={classes.im}>
                    <label htmlFor="image">Movie Image</label>
                    <input type="url" datatype="string" id="image" value={givenImage1} onChange={event => setImage1(event.target.value)}/>
                </div>

                <Stack direction="row" alignItems="center" spacing={2} className={classes.but} >
                    <label htmlFor="contained-button-file">
                        <Input
                            accept="image/*"
                            type="file"
                            id="contained-button-file"
                            value={givenImage2}
                            onChange={event => setImage2(event.target.value)}
                        />
                        <Button variant="contained" component="span">
                            Upload
                        </Button>
                    </label>
                </Stack>

                <div className={classes.control}>
                    <label htmlFor="director">Director's Name</label>
                    <input type="text" required id="director"  datatype="String" value={givenDirector} onChange={event => setDirector(event.target.value)}/>
                </div>

                <div className={classes.control}>
                    <label htmlFor="worth">
                        How Much Do You Think This Movie Is worth Waching?
                    </label>
                    <input type="range" datatype="Int" id="worth" min="1" max="5" value={givenWorth} onChange={event => setWorth(event.target.value)}/>
                </div>

                <div className={classes.control}>
                    <label htmlFor="genre">What is the genre of this movie?</label>
                    <select name="genre" id="genre" value={givenGenre} onChange={event => setGenre(event.target.value)} required datatype="String">
                        <option value="action">Action</option>
                        <option value="drama">Drama</option>
                        <option value="comedy">Comedy</option>
                        <option value="crime">Crime</option>
                        <option value="animation">Animation</option>
                        <option value="fantasy">Fantasy</option>
                        <option value="romance">Romance</option>
                        <option value="thriller">Thriller</option>
                        <option value="horror">Horror</option>
                        <option value="science fiction">Science Fiction</option>
                        <option value="historical">Historical</option>
                        <option value="western">Western</option>
                    </select>
                </div>

                <div className={classes.control}>
                    <label htmlFor="description">Description</label>
                    <textarea
                        id="description"
                        type="text"
                        datatype="String"
                        required
                        rows="5"
                        value={givenDescription} onChange={event => setDescription(event.target.value)}
                    ></textarea>
                </div>

                <div className={classes.control}>
                    <label htmlFor="topic">Review title</label>
                    <textarea
                        id="topic"
                        type="text"
                        datatype="String"
                        required
                        rows="1"
                        value={givenTopic} onChange={event => setTopic(event.target.value)}
                    ></textarea>
                </div>

                <div className={classes.control}>
                    <label htmlFor="review">Review Text</label>
                    <textarea id="review" rows="5" datatype="String" value={givenText} onChange={event => setText(event.target.value)}></textarea>
                </div>

                <div className={classes.ctrl}>
                    <label htmlFor="rank">Add Your Rank</label>
                    <input
                        type="number"
                        name="ranking"
                        id="ranking"
                        min="1"
                        max="100"
                        value={givenRank} onChange={event => setRank(event.target.value)}
                        datatype="Int"
                    ></input>
                </div>

                <div className={classes.actions}>
                    <button type="button" onClick={addMovie}>Add Movie</button>
                </div>
            </form>
        </Card>
    );
}

export default NewMovieForm;