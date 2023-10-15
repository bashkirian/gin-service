import * as React from "react";
import List from '@mui/material/List';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import { Avatar, Button, Divider } from "@mui/material";

import './SearchBox.css'

export default function SearchBox(props) {
    // SelectPosition потом передаётся наверх для передачи в компоненту Map
    // В listPlace заносятся все адреса в listPlace
    const { selectPosition, setSelectPosition, listPlace, setListPlace } = props;

    // Нигде не используется
    const [selectedIndex, setSelectedIndex] = React.useState(1);

    return (
        <div>
            <div className="container">
                <div className="centered-element">
                    <Button
                        variant="contained"
                        color="primary"
                        onClick={() => {
                            // Получаем список всех адресов
                            const fetchData = () => {
                                fetch('/branches')
                                .then(res => res.json())
                                .then(
                                    data => {
                                        setListPlace(data.data);
                                    }
                                )
                            }

                            fetchData();

                            // const fetchData = () => {
                            //     return new Promise((resolve, reject) => {
                            //         fetch('/branches')
                            //             .then(res => res.json())
                            //             .then(
                            //                 data => {
                            //                     setListPlace(data.data);
                            //                 }
                            //             )

                            //         setTimeout(() => {
                            //             const data = "Done";
                            //             resolve(data);
                            //         }, 1000)
                            //     });
                            // }
                            // fetchData().then((data) => {
                            //     // console.log(data);
                            //     // console.log(listPlace);


                            // })

                        }}>
                        Показать на карте
                    </Button>
                </div>
                <div className="right-element">
                    <a href="https://www.vtb.ru/" target="_blank" rel="noreferrer">
                        <Avatar
                            alt=''
                            src='./vtblogo.png'
                            variant="square"
                            sx={{ width: 128, height: 64 }} />
                    </a>
                </div>
            </div>
            {listPlace.map((item) => {
                return (
                    setSelectPosition(item)
                );
            })}
            {/* {listPlace.sort((a, b) => {
                const Lata = a?.Latitude;
                const Lona = a?.Longitude;
                const Latb = b?.Latitude;
                const Lonb = b?.Longitude;

                const pos = [55.8, 37.5]

                var first = Math.sqrt(Math.pow(pos[0] - Lata, 2) + Math.pow(pos[1] - Lona, 2));
                var second = Math.sqrt(Math.pow(pos[0] - Latb, 2) + Math.pow(pos[1] - Lonb, 2));

                if (first < second) {
                    return -1;
                } else if (first > second) {
                    return 1;
                } else {
                    return 0;
                }
            })} */}
            <List component="nav" aria-label="main mailbox folders">
                {/* Выводим адреса
                Каждый адрес имеет кликабельное поле с описанием и иконку */}
                {listPlace.map((item) => {
                    return (
                        <div key={item?.osm_id}>
                            <ListItemButton
                                selected={selectedIndex === 0}
                                onClick={() => {
                                    // console.log(item);
                                    setSelectPosition(item);
                                }
                                }
                            >

                                <ListItemIcon>
                                    <img src="./location.png" alt='' />
                                </ListItemIcon>
                                <ListItemText primary={item?.Name} />
                            </ListItemButton>
                        </div>
                    )
                })}

                <Divider />
            </List>
        </div>
    )
}