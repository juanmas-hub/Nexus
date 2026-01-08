import { CalendarDays, MapPin, Tag } from "lucide-react";
import type { EventItem } from "../types";

interface EventCardProps {
  event: EventItem;
}

const EventCard = ({ event }: EventCardProps) => {
  return (
    <div className="group bg-white rounded-2xl shadow-sm hover:shadow-xl transition-all duration-300 border border-gray-100 overflow-hidden flex flex-col h-full">
      <div className="relative h-48 overflow-hidden">
        <img
          src={event.image}
          alt={event.title}
          className="w-full h-full object-cover transform group-hover:scale-105 transition-transform duration-500"
        />
        <div className="absolute top-4 left-4">
            <span className="inline-flex items-center gap-1 px-3 py-1 text-xs font-medium text-indigo-800 bg-indigo-100/90 backdrop-blur-sm rounded-full">
                <Tag className="w-3 h-3" /> {event.category}
            </span>
        </div>
      </div>

      <div className="p-5 flex-1 flex flex-col justify-between">
        <div>
          <h3 className="text-xl font-bold text-gray-900 mb-2 line-clamp-2 group-hover:text-indigo-600 transition-colors">
            {event.title}
          </h3>

          <div className="space-y-2 mb-4">
            <div className="flex items-center text-sm text-gray-500">
              <CalendarDays className="h-4 w-4 mr-2 text-gray-400" />
              {event.event_date}
            </div>
            <div className="flex items-center text-sm text-gray-500">
              <MapPin className="h-4 w-4 mr-2 text-gray-400" />
              {event.location}
            </div>
          </div>
        </div>

        <div className="flex items-center justify-between pt-4 border-t border-gray-50">
          <div>
            <p className="text-sm text-gray-500">Desde</p>
            <p className="text-2xl font-bold text-gray-900">
              {event.price === 0 ? "Gratis" : `$${event.price}`}
            </p>
          </div>
          <button className="bg-gray-900 text-white px-4 py-2 rounded-lg font-medium hover:bg-indigo-600 transition-colors">
            Reservar
          </button>
        </div>
      </div>
    </div>
  );
};

export default EventCard;