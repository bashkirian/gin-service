import React, { useEffect } from "react";
import { MapContainer, Marker, Popup, TileLayer, useMap } from 'react-leaflet'
import "leaflet/dist/leaflet.css"
import L from "leaflet"

const icon = L.icon({
    iconUrl: "./location.png",
    iconSize: [38, 38]
});

const userIcon = L.icon({
    iconUrl: "./placeholder.png",
    iconSize: [38, 38]
});

const position = [51.505, -0.09]

function ResetCenterView(props) {
    const { selectPosition } = props;
    const map = useMap();

    useEffect(() => {
        if (selectPosition) {
            map.setView(
                L.latLng(selectPosition?.lat, selectPosition?.lon),
                map.getZoom(),
                {
                    animate: true
                }
            )
        }
    }, [selectPosition]);

    return null;
}

export default function Maps(props) {
    const { selectPosition, listPlace } = props;
    // const locationSelection = [selectPosition?.lat, selectPosition?.lon]

    // console.log(listPlace);

    return (
        <MapContainer center={position} zoom={8} style={{ width: '100%', height: '100%' }}>
            <TileLayer
                attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
            />
            {/* Расположение пользователя */}
            <Marker position={position} icon={userIcon}>
                <Popup>
                    User
                </Popup>
            </Marker>

            {/* Расположение банкоматов */}
            {listPlace.map((item) => {
                const locationSelection = [item?.lat, item?.lon];
                // console.log(locationSelection);
                return (
                    selectPosition && (
                        <Marker key={item?.osm_id} position={locationSelection} icon={icon}>
                            <Popup>
                                Marker
                            </Popup>
                        </Marker>
                    )
                )
            })}

            <ResetCenterView selectPosition={selectPosition} />
        </MapContainer>
    )
}