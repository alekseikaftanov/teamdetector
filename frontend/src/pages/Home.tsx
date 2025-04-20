import { useState } from 'react'
import { Logo } from '../components/Logo'
import { LoginModal } from '../components/LoginModal'

const Home = () => {
  const [isLoginModalOpen, setIsLoginModalOpen] = useState(false)

  return (
    <div className="container">
      <header className="header">
        <Logo />
        <nav className="nav">
          <a href="#features" className="nav-link">Преимущества</a>
          <a href="#possibilities" className="nav-link">Возможности</a>
          <a href="#contacts" className="nav-link">Контакты</a>
          <button onClick={() => setIsLoginModalOpen(true)} className="login-button">
            Войти
          </button>
        </nav>
      </header>

      <section className="hero">
        <h1 className="hero-title">
          Полная оценка софт-скилов<br />
          команды за 10 минут
        </h1>
        <p className="hero-subtitle">
          Дадим полную аналитику и рекомендации<br />
          для каждого сотрудника
        </p>
        <div className="hero-buttons">
          <button className="button-primary">Заказать демо</button>
          <button className="button-secondary">Пройти тест</button>
        </div>
      </section>

      <section className="features" id="features">
        <h2 className="features-title">
          <span>Экономим</span>
          <span>🎯</span>
          <span>ваши ресурсы</span>
        </h2>
        <div className="features-grid">
          <div className="feature-card">
            <h3 className="feature-title">Настройте систему для себя</h3>
            <p className="feature-description">
              Гибкие настройки позволяют настроить структуру компании, уровни доступа и критерии оценки
            </p>
          </div>
          <div className="feature-card">
            <h3 className="feature-title">Проверяйте сотрудников на всех уровнях</h3>
            <p className="feature-description">
              Вы можете отправить тест как на всю компанию, так и на отдельные команды, категории сотрудников или соискателей
            </p>
          </div>
          <div className="feature-card">
            <h3 className="feature-title">Простая интерпретация, понятные выводы</h3>
            <p className="feature-description">
              Данные анализируются по разным параметрам, а мы предложим выводы и прогнозы, как игровые механики могут повлиять на вашу компанию
            </p>
          </div>
        </div>
      </section>

      <section className="testimonials">
        <h2 className="section-title">
          <span>С нами растут</span>
          <span>📈</span>
        </h2>
        <div className="testimonials-grid">
          <div className="testimonial-card">
            <p className="testimonial-text">
              «Если раньше критерии оценки софт скилов были субъективными, то сейчас у нас есть понятный и прозрачный инструмент для оценки персонала»
            </p>
            <div className="testimonial-author">
              <span className="author-name">Алексей Кафтанов</span>
              <span className="author-position">Директор компании</span>
            </div>
          </div>
          <div className="testimonial-card">
            <p className="testimonial-text">
              «Если раньше критерии оценки софт скилов были размытыми, и, в основном, субъективными, то сейчас у нас есть понятный и прозрачный инструмент для оценки персонала»
            </p>
            <div className="testimonial-author">
              <span className="author-name">Алексей Кафтанов</span>
              <span className="author-position">Директор компании</span>
            </div>
          </div>
          <div className="testimonial-card">
            <p className="testimonial-text">
              «Если раньше критерии оценки софт скилов были размытыми, и, в основном, субъективными, то сейчас у нас есть понятный и прозрачный инструмент для оценки персонала»
            </p>
            <div className="testimonial-author">
              <span className="author-name">Алексей Кафтанов</span>
              <span className="author-position">Директор компании</span>
            </div>
          </div>
        </div>
      </section>

      <section className="benefits" id="possibilities">
        <h2 className="section-title">
          <span>Лучшее, что вы можете сделать</span>
          <span>🎯</span>
          <span>для своей команды</span>
        </h2>
        <div className="benefits-grid">
          <div className="benefit-card">
            <h3 className="benefit-title">Это не скучно и очень быстро</h3>
            <p className="benefit-description">
              Подпись к заголовку в одну или несколько строк. За головку в одну или несколько строк
            </p>
          </div>
          <div className="benefit-card">
            <h3 className="benefit-title">Понятно, что делать дальше</h3>
            <p className="benefit-description">
              Подпись к заголовку в одну или несколько строк. За головку в одну или несколько строк
            </p>
          </div>
          <div className="benefit-card">
            <h3 className="benefit-title">Понятные критерии оценки</h3>
            <p className="benefit-description">
              Подпись к заголовку в одну или несколько строк. За головку в одну или несколько строк
            </p>
          </div>
          <div className="benefit-card">
            <h3 className="benefit-title">Командная оценка</h3>
            <p className="benefit-description">
              Подпись к заголовку в одну или несколько строк. За головку в одну или несколько строк
            </p>
          </div>
          <div className="benefit-card">
            <h3 className="benefit-title">Конструктор структуры компании</h3>
            <p className="benefit-description">
              Подпись к заголовку в одну или несколько строк. За головку в одну или несколько строк
            </p>
          </div>
        </div>
      </section>

      <section className="cta">
        <div className="cta-content">
          <h2 className="cta-title">Тимдетектед — для команды я выбираю лучшее</h2>
          <p className="cta-description">
            Попробуйте бесплатно, оставьте свои данные и зарезервируйте демо прямо сейчас!
          </p>
          <form className="cta-form">
            <input type="text" placeholder="Почта" className="input" />
            <input type="tel" placeholder="Телефон" className="input" />
            <button type="submit" className="button-secondary">Попробовать бесплатно</button>
          </form>
        </div>
      </section>

      <footer className="footer" id="contacts">
        <div className="footer-content">
          <div className="footer-left">
            <Logo />
            <nav className="footer-nav">
              <a href="#about">О нас</a>
              <a href="#features">Преимущества</a>
              <a href="#possibilities">Возможности</a>
              <a href="#demo">Заказать демо</a>
            </nav>
          </div>
          <div className="footer-right">
            <div className="manager-info">
              <p>А это Максим, наш менеджер,<br />ему можно написать</p>
              <div className="manager-contact">
                <img src="/manager.jpg" alt="Максим Менеджер" className="manager-photo" />
                <span className="manager-name">Максим Менеджер</span>
              </div>
            </div>
          </div>
        </div>
        <div className="footer-bottom">
          <p>© 2023 ТимДетектед. Все права защищены.</p>
        </div>
      </footer>

      {isLoginModalOpen && (
        <LoginModal onClose={() => setIsLoginModalOpen(false)} />
      )}
    </div>
  )
}

export default Home 