import { BACKEND_DOMAIN } from "@/constants/config";
import axios from "axios";
import { RxCross2 } from "react-icons/rx";
import { RxCheck } from "react-icons/rx";
export default function FriendReqeustCard(props: { id: string, tag: string, name: string, email: string, friendRequest: string, acceptCallback: Function, declineCallback: Function }) {
    
    return (
        <div className={`hover:bg-gray-200`}>
            <div className="flex flex-row">
                <div className="m-3">
                    <img className="h-11 w-11 rounded-full" src="data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAkGBwgHBgkIBwgKCgkLDRYPDQwMDRsUFRAWIB0iIiAdHx8kKDQsJCYxJx8fLT0tMTU3Ojo6Iys/RD84QzQ5OjcBCgoKDQwNGg8PGjclHyU3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3N//AABEIAFwAXAMBIgACEQEDEQH/xAAcAAABBQEBAQAAAAAAAAAAAAAFAgMEBgcBAAj/xAA2EAACAQMCAwUGBQQDAQAAAAABAgMABBEFIQYSMSJBUWFxExSBobHBIzJCUpEHM2LRJKLwFf/EABkBAAMBAQEAAAAAAAAAAAAAAAIDBAEABf/EACERAAICAwABBQEAAAAAAAAAAAABAhEDEiExBBQiQVET/9oADAMBAAIRAxEAPwA7XjXa5WDAPxFq50u0BiiMk0mQmThV8yf42rLrktfX0st7LI8jdWG4zV34yja91e2tOYciQc/L4sxIzjyAqZoWiWsIR1iU42DEfOlznqMx49zP/wD5DGIvbhpVB3HLuB50PnjMVwUkUjxGOlb3BpNk8QVreE9dyozvQXWuBbS/lM0MaB8YIzgGgWUbL0/4ZFHbTT/24fTApPusiZLhR4833rb9F4fg02zETxRlx31D1LRbO8VhLChwMZ7xWf26b7fnkonDPEkunp7tKwniADBcdoDvGfKtFjdZI1dfysAR6Vl+qaY2lmWMLzKuSjjYir7wtO1xoFnJIcsUwdqdF2TyVOgrSD1pdJIogSTXKTzVznrTjO+Jrxl4wnTOQoRP+oP3NXbSM+7ovfjJqma/b83HHtOUMHRX6bbLy/UfSrRpN7dGQpb2jMSc8zHGR5VNl6U4HSLbZo5QHpUkBxt9qCxaxJBMIZ7Tt4yRFMjkD0BzRu3u1uIPaxDK+exFJqh+1jMschUtmhs6th8jFOahrMkOSLb8IbGWSRUX51BlvriaAyKLWaM7Ewyh/mK7X7O3+ijccSqsiRgbt137qNcEkHRcK3MqyMBQ3jGIS6X7ZlPPG4wSNxviiHBMEttpZSZChfEgU+B/9mqcb4S5Yu7LDXKVXqaJElsUktnpSTmuCiMAuq2ai6jvsEAOYySdiTvt/FWCwsYbqCMMCVAGEJ2b18aGavHLLps3IexA4kI89vtmiOh3REEXgAKjn5LsaTQYGnWlhac0FvFH0PKi4yaYsXZbecjG7k7VP1J86ZNL1YL2fLzoPYzW3uciTO6yg45GGGJ9KF9GRSQRsY4p1CTBSOXbKgg+Ipu50uzgbnjhiQgHBjTlpOguzuyurgFeZRIhVh8DT2rSAKVOdvOutqNGapyKfr0KXkgtgP7zhfU5AoxcItl7vaY5dm5APDah8MHvOswpzFQhLkjx7vnimJbwzcSi3VuZYEZWb/I4z9BRYvKF5aUWGARXqbFLB86rIxoneoer6gml6bPeyLzCNdlz+ZicAfyRUtqon9RNULSxaXE3ZXEk/mf0j7/EVoLAMfFmsRXlxO1yre8ALJGyZQL5DuwD/vNafocn/GDDHhWM8gJOd60TgzV1lsfYuw9qmEIJ6nuP8fekZY2rHYJtOjQbu8WKJFkYKvQ5pqGe0VeZGj5vErv8KSJhdRpkDY42qUhkQciKpA36ZpCZZ5I638McoaMgH9vQ1y8maVgWBHNvinZvw/xDgnPdQm/1BeeRiQEjGMnvNC+mt0VXjHWrrRp4XsJRHLIxVvEqBv8APFRuDdWsZbqT3icRXT7Ikh/OT1we8/OgvGsj3V1bSuTylGKg+o3qtb7g7+VV40lFEGWT3Zuw60qs44Q4se1dLLU5S9qx5Y5nOTEfAn9v09OmjCmAJhQ8PXPOYxPbGQfoEm/0rBNYna61e9nc5Lzvj0BwPkBX0TqT6VY60b7UdUitijBuWU8gBxt2jtXz1riWqa1fCwlWW1M7mF1OQVJyMHv8K4EhAZGKsXBlr70tzEZDHKpDRyD9Jx8xVfAoxwvqK6fqa+2IWGXCsx/Se40M09eBYmlNWXrStRnsrz3PVVVHkPYkB7D+h8fKrHDqccSEcw5j1PjUO70+HUbMo0at3kEZzQm34ctDKFuDKincYndcEfGpuMs+UQhq2tW0Fqz8+XLdlV3JPkKE6baXd+RcXnYiGSkI8/3f6qaNCs7I89tCCx2DsxZsepzRcRiG2XoDjBArG0vAcU5PpmfHi8uoW6/4t9qqrAHboRV145SJmEjhfaAgJ4+dUlz26oxP4EfqFWRnO/GOtahwrxBZy6HbpeXkcU8I9k4kbBbHQ+e2PjmsxA3zXs0wSmSb6+u9RuWuL+6murhuskzlj8+g8hTYGV5T8K4oFK7q444Njg9aVSW3UHvrqHIB8q4w0T+n/EYcLpN8/wCIBiBz+oftPmO7yq7NCochhlT/ABWEqSrBlJVgcgg4IPiDW08J302rcNWl5eFWnIZWZRjm5WIz67VNljr1FuCey1ZKlTmlQDAx026U1cOMMWYLHGpZmJ2UdaIQwoWbOTgVn/8AU+/uLdbbT4G5ILjnaXHVuUrgZ8N/lSorZ0PlLSNlS4n1VdV1KSWHIt07EWepH7vj9MUDCln26eNPEc3WvN2RtVqSSpHmyk5O2IbbsjrSdhtXer8vdXiMHArTD//Z" />
                </div>
                <div className="flex flex-col justify-center">
                    <span className=" font-medium text-md  text-gray-800">{props.name}</span>
                    <span className=" font-light text-sm  text-gray-500">{props.email}</span>
                </div>
                <div className="flex ml-auto mr-5 justify-center items-center h-auto w-auto">
                    <button className=" p-3 rounded-full active:bg-gray-300 hover:text-green-500 hover:bg-green-100" onClick={()=>props.acceptCallback(props.friendRequest)}><RxCheck className="h-6 w-6" /></button>
                    <button className=" p-3 rounded-full active:bg-gray-300 hover:text-red-500 hover:bg-red-100" onClick={()=>props.declineCallback(props.friendRequest)}><RxCross2 className="h-6 w-6" /></button>

                </div>
            </div>
            <div className="border-t border-gray-300 w-full"></div>
        </div>
    )
}