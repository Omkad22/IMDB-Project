import React from "react";
import {gql, useQuery} from "@apollo/client";
import {Link} from "react-router-dom";

import Button from "@mui/material/Button";
import classes from "./top10.module.css"

import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Typography from '@mui/material/Typography';

function Top10Page() {

    const GET_TOP10_MOVIES = gql`
        query Top10Movies{
            top10Movies {
                title
                rank
                id
                image
            }
        }
    `;

    const {loading, error, data} = useQuery(GET_TOP10_MOVIES)
    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error :</p>;
    let loaded

    console.log(data)

    loaded = data.top10Movies.map(({title, rank, id, image}) => (
        <div>
            <Card sx={{maxWidth: 600}} style={{backgroundColor: "#cc2062", marginBottom: "3cm"}} key={id}>
                <CardMedia
                    component="img"
                    alt="movie image"
                    height="300"
                    src={image}
                />
                <CardContent>
                    <Typography gutterBottom variant="h5" component="div">
                        <p style={{color: "yellow", fontSize: "xx-large"}} className={classes.movie}>
                            {title} : {rank} / 100
                        </p>
                    </Typography>
                </CardContent>
                <CardActions>
                    <Button size="large">Share</Button>
                    <Link to={"/moviePage/" + id} style={{textDecoration: "none"}}><Button size="large">Go To Movie's Page</Button></Link>
                </CardActions>
            </Card>
        </div>
    ))
    return <>{loaded}</>
}

export default Top10Page;
