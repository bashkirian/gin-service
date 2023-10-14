import * as React from "react";
import List from '@mui/material/List';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import { Button, Divider } from "@mui/material";

import './SearchBox.css'

const NOMINATIM_BASE_URL = "https://nominatim.openstreetmap.org/search?q=135+pilkington+avenue,+birmingham&format=xml&polygon_kml=1&addressdetails=1";
// const params = {
//     q: '',
//     format: 'json',
//     addressdetails: 'addressdetails'
// };

export default function SearchBox(props) {
    const { selectPosition, setSelectPosition } = props;

    const [selectedIndex, setSelectedIndex] = React.useState(1);
    const [searchText, setSearchText] = React.useState("");
    const [listPlace, setListPlace] = React.useState([]);

    const handleListItemClick = (
        event: React.MouseEvent<HTMLDivElement, MouseEvent>,
        index: number,
    ) => {
        setSelectedIndex(index);
        console.log('click');
    };

    return (
        <div>
            <Button 
            variant="contained"
            color="primary"
            onClick={() => {
                const params = {
                    q: searchText,
                    format: "json",
                    addressdetails: 1,
                    polygon_geojson: 0,
                };
                const queryString = new URLSearchParams(params).toString();
                const requestOptions = {
                    method: "GET",
                    redirect: "follow",
                };
                fetch(`${NOMINATIM_BASE_URL}${queryString}`, requestOptions)
                .then((response) => response.text())
                .then((result) => {
                    //
                    console.log(JSON.parse(result));
                    //
                    setListPlace(JSON.parse(result));
                })
                .catch((err) => console.log("err: ", err));
            }}>
                Search
            </Button>
            <List component="nav" aria-label="main mailbox folders">
                {listPlace.map((item) => {
                    return (
                        <div key={item?.osm_id}>
                            <ListItemButton
                                // selected={selectedIndex === 0}
                                onClick={(event) => {
                                    setSelectPosition(item);
                                    handleListItemClick(event, 0);
                                }
                            }
                            >
                                <ListItemIcon>
                                    <img src="./location.png" alt='' />
                                </ListItemIcon>
                                <ListItemText primary={item?.display_name} />
                            </ListItemButton>
                        </div>
                    )
                })}
                <Divider />
            </List>
        </div>
    )
}