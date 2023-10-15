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

const position = [55.8, 37.5]
const lockedPosition = [55.419247, 37.570042]

// Центрирование камеры на точке и отрисовка прямой к ней
function ResetCenterView(props) {
    const { selectPosition } = props;
    const map = useMap();

    const pointA = new L.latLng(position[0], position[1]);
    const pointB = new L.latLng(selectPosition?.Latitude, selectPosition?.Longitude);
    const lockedPoint = new L.latLng(lockedPosition[0], lockedPosition[1]);
    const pointList = [pointA, pointB];

    const line = new L.polyline(pointList, {
        color: 'green',
        weight: 3,
        smoothFactor: 1
    });

    useEffect(() => {
        if (selectPosition && (pointB.lat !== lockedPoint.lat)) {
            map.setView(
                L.latLng(selectPosition?.Latitude, selectPosition?.Longitude),
                map.getZoom(),
                {
                    animate: true
                }
            )
            line.addTo(map);
        }
    }, [selectPosition]);

    return null;
}

export default function Maps(props) {
    const { selectPosition, listPlace } = props;

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
                const locationSelection = [item?.Latitude, item?.Longitude];
                return (
                    selectPosition && (
                        <Marker key={item?.osm_id} position={locationSelection} icon={icon}>
                            <Popup>
                                {item?.Address}
                            </Popup>
                        </Marker>
                    )
                )
            })}

            <ResetCenterView selectPosition={selectPosition} />
        </MapContainer>
    )
}