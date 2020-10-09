import {
  Card,
  CardContent,
  CardMedia,
  Grid,
  makeStyles,
  Typography,
} from "@material-ui/core";
import React from "react";
import { withRouter } from "react-router";

import logo from "../assets/logo.png";

const useStyles = makeStyles(() => ({
  container: {
    margin: 10,
  },
  gridItem: {
    padding: 10,
  },
  media: {
    height: 0,
    paddingTop: "100%", // 16:9
  },
}));

/*
data = [{
        version: string
        last_updated: datetime-string
    }...
]
*/
export const RepoView = withRouter(function ({ data, match }) {
  const releases = data;
  const repoName = match.params.name;

  const styles = useStyles();

  return (
    <Grid container className={styles.container}>
      <Grid item xs={3} className={styles.gridItem}>
        <Card>
          <CardMedia
            image={logo}
            title="Repository name comes here"
            className={styles.media}
          />
          <CardContent>
            <Typography variant="h5">{repoName}</Typography>
          </CardContent>
        </Card>
      </Grid>
      <Grid item xs={9} className={styles.gridItem}>
        Hello world
      </Grid>
    </Grid>
  );
});
