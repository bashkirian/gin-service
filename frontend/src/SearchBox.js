import * as React from "react";
import List from '@mui/material/List';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import { Avatar, Button, Divider, Icon, SvgIcon } from "@mui/material";
import axios from 'axios'

import './SearchBox.css'

const NOMINATIM_BASE_URL = "https://nominatim.openstreetmap.org/search?q=135+pilkington+avenue,+birmingham&format=xml&polygon_kml=1&addressdetails=1";
// const params = {
//     q: '',
//     format: 'json',
//     addressdetails: 'addressdetails'
// };
// const fileJson = "./atms.json";

export default function SearchBox(props) {
    // SelectPosition потом передаётся наверх для передачи в компоненту Map
    // В listPlace заносятся все адреса в listPlace
    const { selectPosition, setSelectPosition, listPlace, setListPlace } = props;

    // Нигде не используется
    const [selectedIndex, setSelectedIndex] = React.useState(1);
    // Формально используется при принятии jsonа
    const [searchText, setSearchText] = React.useState("");

    // Используется для отображения 'Загрузка...'
    // const [isLoaded, setIsLoaded] = React.useState(false);
    // Заносятся все адреса в listPlace
    // const [listPlace, setListPlace] = React.useState([]);

    // const handleListItemClick = (
    //     event: React.MouseEvent<HTMLDivElement, MouseEvent>,
    //     index: number,
    // ) => {
    //     setSelectedIndex(index);
    //     console.log('click');
    // };

    return (
        <div>
            <div className="container">
                <div className="centered-element">
                    <Button
                        variant="contained"
                        color="primary"
                        onClick={() => {
                            // const params = {
                            //     q: searchText,
                            //     format: "json",
                            //     addressdetails: 1,
                            //     polygon_geojson: 0,
                            // };
                            // const queryString = new URLSearchParams(params).toString();
                            // const requestOptions = {
                            //     method: "GET",
                            //     redirect: "follow",
                            // };

                            // fetch(`${NOMINATIM_BASE_URL}${queryString}`, requestOptions)
                            //     .then((response) => response.text())
                            //     .then((result) => {
                            //         // setIsLoaded(true);
                            //         //
                            //         console.log(JSON.parse(result));
                            //         //
                            //         setListPlace(JSON.parse(result));
                            //     })
                            //     .catch((err) => {
                            //         // setIsLoaded(true);
                            //         console.log("err: ", err);
                            //     });

                            const fetchData = () => {
                                fetch('/branches')
                                .then(res => res.json())
                                .then(
                                    data => {
                                        setListPlace(data.data);
                                        console.log(data.data);
                                    }
                                )
                            }

                            fetchData()
                        }}>
                        Показать на карте
                    </Button>
                </div>
                <div className="right-element">
                    <a href="https://www.vtb.ru/" target="_blank">
                        <Avatar 
                        alt='' 
                        src='./vtblogo.png' 
                        variant="square" 
                        sx={{ width: 72, height: 72 }}/>
                    </a>
                </div>
            </div>
            <List component="nav" aria-label="main mailbox folders">
                {/* Выводим адреса
                Каждый адрес имеет кликабельное поле с описанием и иконку */}
                {listPlace.map((item) => {
                    return (
                        <div key={item?.osm_id}>
                            <ListItemButton
                                selected={selectedIndex === 0}
                                onClick={(event) => {
                                    setSelectPosition(item);
                                    // handleListItemClick(event, 0);
                                }
                                }
                            >

                                <ListItemIcon>
                                    <img src="./location.png" alt='' />
                                </ListItemIcon>
                                <ListItemText primary={item?.salePointName} />
                            </ListItemButton>
                        </div>
                    )
                })}

                <Divider />
            </List>
        </div>
    )
}