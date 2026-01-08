import { useEffect, useState } from "react";
import Navbar from "./components/Navbar";
import Hero from "./sections/Hero";
import EventCard from "./components/EventCard";
import LoginForm from "./components/LoginForm";
import RegisterForm from "./components/RegisterForm";
import { ArrowRight, Loader2 } from "lucide-react";
import type { EventItem } from "./types"; 
import { useAuth } from "./context/AuthContext";
import { getEvents } from "./api/eventsService";

function App() {
  const { isAuthenticated } = useAuth();
  const [authView, setAuthView] = useState<"login" | "register">("login");
  
  const [events, setEvents] = useState<EventItem[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (isAuthenticated) {
      setIsLoading(true);
      getEvents()
        .then((data) => {
          setEvents(data);
          setError(null);
        })
        .catch((err) => {
          console.error("Error en Nexus Gateway:", err);
          setError("No se pudieron cargar los eventos. Verificá que el servicio de Rust esté online.");
        })
        .finally(() => setIsLoading(false));
    }
  }, [isAuthenticated]);

  if (!isAuthenticated) {
    return authView === "login" ? (
      <LoginForm onSwitchRegister={() => setAuthView("register")} />
    ) : (
      <RegisterForm onSwitchLogin={() => setAuthView("login")} />
    );
  }

  return (
    <div className="bg-gray-50 min-h-screen font-sans">
      <Navbar />
      <main>
        <Hero />
        <section className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16">
          <div className="flex justify-between items-end mb-10">
            <div>
              <h2 className="text-3xl font-bold text-gray-900 mb-2">Próximos Eventos</h2>
              <p className="text-gray-600 text-lg">Explora lo que está pasando en Nexus.</p>
            </div>
            <button className="hidden md:flex items-center text-indigo-600 hover:text-indigo-700 font-semibold transition group">
              Ver todos <ArrowRight className="ml-2 h-5 w-5 group-hover:translate-x-1 transition-transform"/>
            </button>
          </div>

          {isLoading ? (
            <div className="flex flex-col items-center justify-center py-20">
              <Loader2 className="h-12 w-12 text-indigo-600 animate-spin mb-4" />
              <p className="text-gray-500 italic">Conectando con el Catalog Service (Rust)...</p>
            </div>
          ) : error ? (
            <div className="bg-red-50 border border-red-200 text-red-700 px-6 py-4 rounded-lg text-center">
              {error}
            </div>
          ) : (
            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6 lg:gap-8">
              {events.map((event: EventItem) => (
                <EventCard key={event.id} event={event} />
              ))}
            </div>
          )}

          <div className="mt-12 flex justify-center md:hidden">
             <button className="w-full sm:w-auto flex items-center justify-center bg-white border border-gray-300 px-6 py-3 rounded-full text-gray-700 font-semibold hover:bg-gray-50 transition shadow-sm">
                Ver todos los eventos <ArrowRight className="ml-2 h-5 w-5"/>
             </button>
          </div>
        </section>
      </main>
      <footer className="bg-gray-900 text-white py-12 text-center">
          <p className="text-gray-400">© 2026 Nexus Events. Desarrollado en Go y Rust.</p>
      </footer>
    </div>
  );
}

export default App;